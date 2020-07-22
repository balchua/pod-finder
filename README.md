# Pod finder

A simple go application which finds all the pods based on a label and prints to a file the pod names and IP addresses.

The generated file will look like this

```json
{
  "pods": [
    {
      "name": "mypod-54f75566fb-6vmbv",
      "ip": "10.1.81.30",
      "status": "Running"
    },
    {
      "name": "mypod-54f75566fb-3hgt5",
      "ip": "10.1.81.31",
      "status": "Running"
    }
  ]
}
```

Empty data

```json
{
  "pods": []
}
```

## Running the application

This section shows how to run the application in either inside a kubernetes cluster or outside the kubernetes cluster

### In cluster


### Out of cluster

Make sure you have the kube config file.  
The example below shows how to run and find all pods:

* With the label `app: activemq-artemis` from the artemis `namespace`.
* With the location of the kube config file taken from `$KUBECONFIG`
* With `period` set to `5` seconds
* With the `output` written to the file `/tmp/output.json`

```shell
$ ./pod-finder check --config $KUBECONFIG --namespace artemis --period 5 --output /tmp/output.json --label app=activemq-artemis

INFO[0000] checking pods in namespace [artemis]         
INFO[0000] Using out of cluster config                  
INFO[0001] Start retrieving pods in namespace artemis   
INFO[0001] Finished retrieving pods in namespace artemis 
INFO[0002] Start retrieving pods in namespace artemis   
INFO[0002] Finished retrieving pods in namespace artemis 
INFO[0003] Start retrieving pods in namespace artemis   
INFO[0003] Finished retrieving pods in namespace artemis 
INFO[0004] Start retrieving pods in namespace artemis   
INFO[0004] Finished retrieving pods in namespace artemis 
INFO[0005] Start retrieving pods in namespace artemis   
INFO[0005] Finished retrieving pods in namespace artemis 
^CINFO[0005] Got interrupt signal. Aborting...            
INFO[0005] One last run before shutting down.           
INFO[0005] Start retrieving pods in namespace artemis   
INFO[0005] Finished retrieving pods in namespace artemis 
```

## Build

The project uses skaffold to build and run the application in container.

To simply build the image.

Change the following in the [`skaffold.yaml`](skaffold.yaml)
`$ skaffold build`
