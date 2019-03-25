package main

import (
	"github.com/shinytang6/openpaas-security/client"
	"github.com/shinytang6/openpaas-security/deployment"
)

func main() {
	clientset := client.InitClient()
	deploymentClient := clientset.AppsV1beta1().Deployments("default")
	//fmt.Println(deployment)
	//fmt.Println(deployment.Items[0].Spec.Template.Spec.Containers)
	//fmt.Println(len(deployment.Items))
	//fmt.Println(clientset)
	deployment.CreateDeployment(deploymentClient)
}