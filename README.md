# time
4p : 4min
8p : 6min15




10p : 7min43
19p : 13min

# 事前準備
- [ ] ArgoCD port-forward
  - [ ] `kubectl port-forward svc/argocd-server -n argocd 8080:443`
  - [ ] initial password
    - [ ] `kubectl -n argocd get secret/argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo` 
    - [ ] 2KG7OcaBzAKswWrQ
  - [ ] http://localhowst:8080
- [ ] Argo Rollout
  - [ ] `kubectl argo rollouts dashboard`
  - [ ] http://localhowst:3100
- [ ] Demo1, Demo2とArtifact Registry と Githubの画面分割を用意しておく
- [ ] ターミナル開いておく（cat でスクリプトの中も見れるようにしておく）


# Demo1


# Demo2

```bash
# 3秒ごとにリクエストを送信するシェルスクリプトを実行
$ sh access.sh

# 別ターミナルで確認
$ tail -f log.txt
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