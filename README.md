# Movie_Data_Rename

## build

```bash
docker run \
--rm \
-v $PWD/:/go/src/github.com/mengkzhaoyun/movie_data_rename \
-w /go/src/github.com/mengkzhaoyun/movie_data_rename \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.21-alpine \
bash .beagle/build.sh
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="Mengkzhaoyun" \
  -e DRONE_REPO_NAME="Movie_Data_Rename" \
  -e PLUGIN_MOUNT="./.git" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="Mengkzhaoyun" \
  -e DRONE_REPO_NAME="Movie_Data_Rename" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```

## debug

```bash
docker run -it --rm \
-v $PWD/.tmp:/data \
registry.cn-qingdao.aliyuncs.com/wod/movie_data_rename:v1.0.0
```
