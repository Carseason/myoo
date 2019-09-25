#!/bin/bash
kill -9 $(ps -ef | grep myoo | grep -v grep | awk '{print $2}');
wget -O myoo https://github.com/Carseason/myoo/raw/master/myoo;
wget -O src/js/app.js https://github.com/Carseason/myoo/raw/master/src/js/app.js;
wget -O src/js/style.css https://github.com/Carseason/myoo/raw/master/src/js/style.css;
chmod 0755 myoo;
nohup ./myoo &