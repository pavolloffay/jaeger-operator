package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jaegertracing/jaeger-operator/pkg/apis/jaegertracing/v1"
	"github.com/jaegertracing/jaeger-operator/pkg/util"
)

const (
	workingDir = "/tmp/_working_dir"
	certScript = "./scripts/cert_generation.sh"
)

type secret struct {
	name    string
	content map[string]string
}

var esSecret = secret{
	name: "elasticsearch",
	content: map[string]string{
		"elasticsearch.key": "elasticsearch.key",
		"elasticsearch.crt": "elasticsearch.crt",
		"logging-es.key":    "logging-es.key",
		"logging-es.crt":    "logging-es.crt",
		"admin-key":         "system.admin.key",
		"admin-cert":        "system.admin.crt",
		"admin-ca":          "ca.crt",
	},
}

var jaegerSecret = secret{
	name: "jaeger-elasticsearch",
	content: map[string]string{
		"ca": "ca.crt",
	},
}

var curatorSecret = secret{
	name: "curator",
	content: map[string]string{
		"ca":   "ca.crt",
		"key":  "system.logging.curator.key",
		"cert": "system.logging.curator.crt",
	},
}

func secretName(jaeger, secret string) string {
	return fmt.Sprintf("%s-%s", jaeger, secret)
}

// ESSecrets assembles a set of secrets related to Elasticsearch
func ESSecrets(jaeger *v1.Jaeger) []corev1.Secret {
	return []corev1.Secret{
		// master and ES secrets use hardcoded name - e.g. do not use instance name in it
		// the other problem for us is that sg_config.yml defines a role which depends on namespace
		// we could make the "resource" configurable once ES image and es-operator-are refactored
		// https://jira.coreos.com/browse/LOG-326
		createSecret(jaeger, esSecret.name, getWorkingDirContents(esSecret.content)),
		createSecret(jaeger, secretName(jaeger.Name, jaegerSecret.name), getWorkingDirContents(jaegerSecret.content)),
		createSecret(jaeger, secretName(jaeger.Name, curatorSecret.name), getWorkingDirContents(curatorSecret.content)),
	}
}

// CreateESCerts runs bash scripts which generates certificates
func CreateESCerts() error {
	return createESCerts(certScript)
}

// createESCerts runs bash scripts which generates certificates
func createESCerts(script string) error {
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		return fmt.Errorf("failed to get watch namespace: %v", err)
	}
	// #nosec   G204: Subprocess launching should be audited
	cmd := exec.Command("bash", script)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.WithFields(log.Fields{
			"script": script,
			"out":    string(out)}).
			Error("Failed to create certificates")
		return fmt.Errorf("error running script %s: %v", script, err)
	}
	return nil
}

func createSecret(jaeger *v1.Jaeger, secretName string, data map[string][]byte) corev1.Secret {
	return corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: jaeger.Namespace,
			Labels: map[string]string{
				"app":                          "jaeger",
				"app.kubernetes.io/name":       secretName,
				"app.kubernetes.io/instance":   jaeger.Name,
				"app.kubernetes.io/component":  "es-secret",
				"app.kubernetes.io/part-of":    "jaeger",
				"app.kubernetes.io/managed-by": "jaeger-operator",
			},
			OwnerReferences: []metav1.OwnerReference{util.AsOwner(jaeger)},
		},
		Type: corev1.SecretTypeOpaque,
		Data: data,
	}
}

func getWorkingDirContents(content map[string]string) map[string][]byte {
	c := map[string][]byte{}
	for secretKey, certName := range content {
		c[secretKey] = getWorkingDirFileContents(certName)
	}
	return c
}

func getWorkingDirFileContents(filePath string) []byte {
	return getFileContents(getWorkingDirFilePath(filePath))
}

func getWorkingDirFilePath(toFile string) string {
	return path.Join(workingDir, toFile)
}

func getFileContents(path string) []byte {
	if path == "" {
		return nil
	}
	contents, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil
	}
	return contents
}
