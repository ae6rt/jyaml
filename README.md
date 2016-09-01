## jyaml - a tool for converting YAML to JSON and vice versa

Read standard input and convert to or from JSON / YAML.  jyaml autosenses the input to produce the complementary output.

### Build

```
go get github.com/ghodss/yaml
go build
```

### Usage

Given an input file in YAML

```
$ cat pod.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: dnsmasq
spec:
  containers:
  - env:
    - name: MY_POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    image: docker-registry.dev.xoom.com/markp/dnsmasq:2.76
    imagePullPolicy: Always
    name: dnsmasq
    ports:
    - containerPort: 53
      name: server-port
      protocol:  UDP
  restartPolicy: Always
```

convert it to JSON:

```
$ ./jyaml < pod.yaml 
{"apiVersion":"v1","kind":"Pod","metadata":{"name":"dnsmasq"},"spec":{"containers":[{"env":[{"name":"MY_POD_NAME","valueFrom":{"fieldRef":{"fieldPath":"metadata.name"}}}],"image":"docker-registry.dev.xoom.com/markp/dnsmasq:2.76","imagePullPolicy":"Always","name":"dnsmasq","ports":[{"containerPort":53,"name":"server-port","protocol":"UDP"}]}],"restartPolicy":"Always"}}
```

and back to YAML

```
$ ./jyaml < pod.yaml | ./jyaml 
apiVersion: v1
kind: Pod
metadata:
  name: dnsmasq
spec:
  containers:
  - env:
    - name: MY_POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    image: docker-registry.dev.xoom.com/markp/dnsmasq:2.76
    imagePullPolicy: Always
    name: dnsmasq
    ports:
    - containerPort: 53
      name: server-port
      protocol: UDP
  restartPolicy: Always
```
