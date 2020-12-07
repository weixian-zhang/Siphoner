package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"
	//apimv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/labels"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPods() {
	//https://stackoverflow.com/questions/51106923/labelselector-for-secrets-list-in-k8s
	
	// labelSelectors := apimv1.LabelSelector{MatchLabels: labels}
	// podListOpts := apimv1.GetOptions {
	// 	LabelSelector: labelSelectors,
	// }

	// pod :=  v1.Pod{}
	// pod.Ge
}

func GetPodLogs(namespace string, name string, container string) (error) {

	clientset, err := newClientSet()
	
	Stdlog.Err(err)

	podLogOptions := v1.PodLogOptions{
		Container: container,
	}

	podLogReq := clientset.CoreV1().Pods(namespace).GetLogs(name, &podLogOptions)

	ioReadCloser, err := podLogReq.Stream()
	Stdlog.Err(err)
	defer ioReadCloser.Close()

	for {

		buf := make([]byte, 2000)
		numBytes, err := ioReadCloser.Read(buf)
		Stdlog.Err(err)

		if(numBytes == 0) {
			break
		}

		if(err == io.EOF) {
			break
		}

		if err != nil {
            return err
		}
		
		message := string(buf[:numBytes])

		Stdlog.Info(message)
	}

	return err

}

func newClientSet() (*kubernetes.Clientset, error) {

	homeDir := homeDir()
	kubeconfig := flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "kubeconfig file")
	//flag.String("kubeconfig", "~/.kube/config", "kubeconfig file") //linux dev

	flag.Parse()

	var kubeConfig *rest.Config

	outClusterConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	Stdlog.Err(err)

	if err != nil {

		inClusterConfig, err := rest.InClusterConfig()
		Stdlog.Err(err)

		if err != nil {
			return nil, err
		}

		kubeConfig = inClusterConfig
	}

	kubeConfig = outClusterConfig

	clientset, err := kubernetes.NewForConfig(kubeConfig)

	return clientset, err
}

func getConfigFromConfigMap(namespace string, name string) {

	_, err := newClientSet()
	Stdlog.Err(err)

	
	//configMap, err := kubeclient.CoreV1().ConfigMaps(namespace).Get(name, nil)
}


func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}