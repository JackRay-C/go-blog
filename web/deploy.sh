#! /bin/bash


echo "编译..."
yarn build 

deploy_path = /opt/go_blog/www/

echo "删除旧版本"
ssh root@106.12.146.78 "rm -rf "${deploy_path}

echo "上传新版本"
scp -r ./dist  root@106.12.146.78:${deploy_path}

echo "重启nginx"
ssh root@106.12.146.78 "nginx -s reload"

echo "部署成功."