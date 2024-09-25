# gitops-demo

```bash
# build コマンド
$ gcloud builds submit --tag asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo:test .


# build
$ docker build -t asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo/memegen:blue . --no-cache

# tag
$ docker tag asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo/memegen:blue memegen:blue

# push
$ docker push asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo/memegen:blue


- 事前にbuild済みimageをArtifactRegistryにtagをつけてpushしておく
- Github上でtagを変更して、refresh → Syncを行う
# ArgoCD
port-forward
- kubectl port-forward svc/argocd-server -n argocd 8080:443
- http://localhost:8080

初期パスワード
- kubectl -n argocd get secret/argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo 