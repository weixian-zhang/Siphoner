package kube

import (
	"flag"
	"io"
	"os"
	"path/filepath"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/weixian-zhang/Siphoner/main"
)

var kubeclient *kubernetes.Clientset

func GetPodsByFilteredNamespaces(namespaces []string, podLabels map[string]string) ([]PodInfo, error) {
	//https://stackoverflow.com/questions/51106923/labelselector-for-secrets-list-in-k8s

	var pods []main.PodInfo
	
	//filter pods by labels if specified
	labelSelector := metav1.LabelSelector{MatchLabels: podLabels}
	podListOpts := metav1.ListOptions {
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	}

	for _, ns := range namespaces {
		podL, err := kubeclient.CoreV1().Pods(ns).List(podListOpts)
		main.Stdlog.Err(err)

		if err != nil {
			return pods, err
		}

		for _, p := range podL.Items {

			podInfo := main.PodInfo{
				Namespace: p.Namespace,
				Name: p.Name,
				Labels: p.GetObjectMeta().GetLabels(),
			}

			for _, v := range p.Spec.Containers {
				podInfo.ContainerNames = append(podInfo.ContainerNames, v.Name)
			}

			pods = append(pods, podInfo)
		}
	}

	return pods, nil
}

func getPodLogs(pods []main.PodInfo) {

	for _, p := range pods {
		for _, c := range p.ContainerNames {

			
			GetContainerLogs(p.Namespace, p.Name, c)

		} 
	}
}

func GetContainerLogs(namespace string, podName string, container string) (error) {
	
	// kubeTime := metav1.Time{Time: time.Now()}
	// kubeTime.Add(-50 * time.Second)

	getLogSinceTime := metav1.Time{Time: time.Now().Add(-3600 * time.Second)}

	podLogOptions := v1.PodLogOptions{
		Container: container,
		Timestamps: true,
		SinceTime: &getLogSinceTime,
	}

	//https://stackoverflow.com/questions/47915287/where-are-kubernetes-pods-logfiles
	//https://kubernetes.io/docs/concepts/cluster-administration/logging/
	
	podLogReq := kubeclient.CoreV1().Pods(namespace).GetLogs(podName, &podLogOptions)

	ioReadCloser, err := podLogReq.Stream()

	main.Stdlog.Err(err)
	if err != nil {
		return err
	}

	defer ioReadCloser.Close()

	for {

		buf := make([]byte, 2000)
		numBytes, err := ioReadCloser.Read(buf)
		main.Stdlog.Err(err)

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

		main.Stdlog.Info(message)
	}

	return err
}

func initKubeClientSet() (error) {

	homeDir := homeDir()
	kubeconfig := flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "kubeconfig file")
	//flag.String("kubeconfig", "~/.kube/config", "kubeconfig file") //linux dev

	flag.Parse()

	var kubeConfig *rest.Config

	outClusterConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	main.Stdlog.Err(err)

	if err != nil {

		inClusterConfig, err := rest.InClusterConfig()
		main.Stdlog.Err(err)

		if err != nil {
			return err
		}

		kubeConfig = inClusterConfig
	}

	kubeConfig = outClusterConfig

	clientset, err := kubernetes.NewForConfig(kubeConfig)

	if err == nil {
		kubeclient = clientset
	}

	return err
}

func initConfigFromConfigMapList() (v1.ConfigMap) {
	
	cmList, err := kubeclient.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
	
	if err != nil {
		main.Stdlog.Err(err)
	}

	var cm v1.ConfigMap = v1.ConfigMap{}

	for _, v := range cmList.Items {
		labelVal := v.Labels[main.ConfigMapLabelKey]
		if labelVal == main.ConfigMapLabelVal {
			cm = v
			break
		}
	}

	return cm
}

func watchConfigMapChanges() {
	factory := informers.NewSharedInformerFactory(kubeclient, 0)

	informer := factory.Core().V1().ConfigMaps().Informer()

	stopWatching := make(chan struct{})
	defer close(stopWatching)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
		   cm, _ := obj.(v1.ConfigMap)
		   main.onConfigAdded(cm)
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			cm, _ := newObj.(v1.ConfigMap)
			main.onConfigUpdated(cm)
		},
	})

}

func getTerminusSecrets() {

}


func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}