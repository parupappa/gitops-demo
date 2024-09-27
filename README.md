
# 事前準備
- [ ] ArgoCD port-forward
  - [ ] `kubectl port-forward svc/argocd-server -n argocd 8080:443`
  - [ ] initial password
    - [ ] `kubectl -n argocd get secret/argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo` 
  - [ ] http://localhowst:8080
- [ ] Argo Rollout
  - [ ] `kubectl argo rollouts dashboard`
  - [ ] http://localhowst:3100


# Demo1


# Demo2

```bash
# 3秒ごとにリクエストを送信するシェルスクリプトを実行
$ sh access.sh

# 別ターミナルで確認
$ tail -f ./gitops-demo/weather-app/log.txt
```

# Commnad
```bash
# Set up the environment
$ export PROJECT_ID=test-yokoo

# build & push ArtifactRegistry
$ gcloud builds submit \
  --project test-yokoo \    
  --tag asia-northeast1-docker.pkg.dev/test-yokoo/$PROJECT_ID/memegen:green \
  --region asia-northeast1 \
  .
```