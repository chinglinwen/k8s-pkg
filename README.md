# k8s
some k8s customization packages

## Endpoint 

Since package import path direct change will need vendor dependency too

Here, we just create symbol link to the target place, replace the origin one

```
cd /home/wen/gocode/src/k8s.io/kubernetes/pkg/controller
mv endpoint{,.bak}
ln -s /home/wen/gocode/src/github.com/chinglinwen/k8s/pkg/controller/endpoint .
```
