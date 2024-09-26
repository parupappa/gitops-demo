
# 事前準備
- [ ] ArgoCDのport-forwardを行う
  - [ ] `kubectl port-forward svc/argocd-server -n argocd 8080:443`
- [ ] チェックボックス1

## local
```bash
# Set up the environment
$ export PROJECT_ID=test-yokoo

# build
$ docker build -t asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue . --no-cache

# tag
$ docker tag asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue memegen:blue

# push
$ docker push asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue
```

## On Cloud Build
```bash
# Set up the environment
$ export PROJECT_ID=test-yokoo

# build & push ArtifactRegistry
# memegen 配下で実行すること
$ gcloud builds submit \
  --project test-yokoo \    
  --tag asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo/memegen:green \
  --region asia-northeast1 \
  .
```

- 事前にbuild済みimageをArtifactRegistryにtagをつけてpushしておく
- Github上でtagを変更して、refresh → Syncを行う
# ArgoCD
port-forward
- kubectl port-forward svc/argocd-server -n argocd 8080:443
- http://localhost:8080

初期パスワード
- kubectl -n argocd get secret/argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo 


# Demo1


# Demo2

```bash
# 3秒ごとにリクエストを送信するシェルスクリプトを実行
$ sh access.sh

# 別ターミナルで確認
$ tail -f ./gitops-demo/weather-app/log.txt
```