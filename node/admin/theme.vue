<template>
    <div class="body">
        <div class="theme">
            <li class="theme-list">
                <div class="theme-f">主题信息:</div>
                <a class="theme-r">Myoo - 好想告诉你</a>
            </li>
            <li class="theme-list">
                <div class="theme-f">主题版本:</div>
                <a class="theme-r">v1.0</a>
            </li>
            <li class="theme-list">
                <div class="theme-f">主题作者:</div>
                <a class="theme-r">Carseason</a>
            </li>
            <li class="theme-list">
                <div class="theme-f">发布地址:</div>
                <a href="https://github.com/Carseason/myoo" target="_blank" class="theme-r">https://github.com/Carseason/myoo</a>
            </li>
            <li class="theme-list">
                <div class="theme-f"></div>
                <a class="theme-r"></a>
            </li>

            <!-- <div class="receipt">
                <img src="https://i.loli.net/2019/09/09/VbWJMk3vCNFwacI.jpg" alt />
                <img src="https://i.loli.net/2019/09/09/yU6hRuG1sq8z4wn.jpg" alt />
            </div>-->
        </div>
    </div>
</template>
<script>
export default {

}
</script>
<style scoped>
.theme {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 500px;
    margin: 0 auto;
    text-align: center;
    color: #fff;
    display: flex;
    flex-wrap: wrap;
}
li.theme-list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: flex;
    flex-wrap: wrap;
    margin: 10px 0;
    font-size: 1.5rem;
}

.theme-f {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100px;
}

.theme-r {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 400px;
}
.theme-r {
    color: #fff;
}
.theme-r:hover {
    color: rgb(255, 159, 198);
}
</style>
