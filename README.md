# Pod finder

A simple go application which finds all the pods based on a label and prints to a file the pod names and IP addresses.

The generated file will look like this

```json
{
    [
        {
            "name": "pod-name",
            "ip": "pod-ip",
            "status": "status"
        },
        {
            "name": "pod-name",
            "ip": "pod-ip",
            "status": "status"
        }
    ]
}
```

## Running the application

This section shows how to run the application in either inside a kubernetes cluster or outside the kubernetes cluster

### In cluster


### Out of cluster


## Build