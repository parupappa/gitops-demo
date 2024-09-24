# gitops-demo

```bash
# build コマンド
$ gcloud builds submit --tag asia-northeast1-docker.pkg.dev/test-yokoo/gitops-demo:test .
```



- 事前にbuild済みimageをArtifactRegistryにtagをつけてpushしておく
- Github上でtagを変更して、refresh → Syncを行う