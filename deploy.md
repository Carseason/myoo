#   部署教程

#   部署步骤
    1.下载 conf/config src/* views/index.html myoo
    2.配置 conf/config.json文件,
        Port 为程序端口,格式应为 :8080
        数据库默认使用mysql,请在 mysql 下创建 要使用的表 如 myoo 
            Db.Uname 数据库账号
            DB.Pass  数据库密码
            Db.Table 数据库表名称 
            Db.Addr  数据库地址,如 127.0.0.1:3306
        默认使用 Redis 作为缓存， 可不使用
            Redis.Addr  Redis地址,如 127.0.0.1:6379
            Redis.Pass Redis链接密码
        Key ---重要，作用于 jwt 加密 以及 程序对数据加密的密钥,请使用 16 位的密钥
    

#   独立部署
    1.修改 conf/config.json 里的 ServerRender 为 true
    2.使用 nginx 代理端口。nginx需要添加如下规则:
        location / {
            try_files /_not_exists_ @backend;
        }
        location @backend {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;
            proxy_pass http://127.0.0.1:8080;
        }

#   前后端分离部署
    1.修改 conf/config.json 里的 ServerRender 为 false
    2.修改 views/index.html 里的 <meta name="server" id="myoo-server" content="后端域名,结尾不带/">
    3.修改 前端 nginx 配置文件，添加如下项:
        location ^~/src {
            alias /程序文件目录/src/;
        }
        location / {
            try_files $uri $uri/ /index.html;
        }
    4.修改 后端 nginx 配置文件,添加如下项:
        location / {
            try_files /_not_exists_ @backend;
            add_header Access-Control-Allow-Origin *;
            add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
            add_header Access-Control-Allow-Headers 'X-MYOO';
            if ($request_method = 'OPTIONS') {
                return 204;
            }
        }
        location @backend {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;
            proxy_pass http://127.0.0.1:8080;
        }

#   应用程序启动
    1.在程序目录下执行 ./myoo 即可,请至少要有可执行权限
    2.程序会在你数据库下创建如下表:
            category    //分类
            comments    //评论
            favorite    //文章收藏
            followers   //关注粉丝
            options     //配置信息
            posts       //文章
            posts_meta  //文章资源,如视频 音乐 下载信息
            posts_supported //文章点赞
            user        //用户
    3.程序启动完毕后,进入到注册页面 /sign/registered 注册用户，
        注册成功后需要把数据库 user 表的注册用户的 admin 字段改为 1 ,意为管理员，
        如果你配置了redis缓存,则需要 进入到 redis-cli 里 执行 del Myoo/User/1 命令删除当前用户缓存。
    4.配置网站信息, 进入 /admin 后台进行配置以及添加用户组,配置完成后可点击 删除缓存。