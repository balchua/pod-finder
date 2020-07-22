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

## Running the application

This section shows how to run the application in either inside a kubernetes cluster or outside the kubernetes cluster

### In cluster


### Out of cluster


## Build