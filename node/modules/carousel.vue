<template>
    <div class="myoo-carousel">
        <div class="carousel-content" :style="'margin-left:-' + Status + '00%'">
            <template v-for="(value, i) in carousel">
                <li :class="{on : Status == i}" :key="i">
                    <a :href="$root.GetPostsUrl(value.Id)" target="_blank">
                        <img v-lazy="value.Thumbnail" :key="value.Id" />
                        <div class="text" :title="value.Title">{{value.Title}}</div>
                    </a>
                </li>
            </template>
        </div>
        <div class="carousel-trig">
            <template v-for="(value,i) in Len">
                <li :class="{on : Status == i }" v-on:click="SetHdp(i)" :key="i"></li>
            </template>
        </div>
    </div>
</template>

<script>
export default {
    props: ['carousel'],
    data: function () {
        return {
            Time: 5000,
            Len: 0,
            Status: -1,
        }
    },
    mounted: function () {
        this.Len = this.carousel.length;
        this.HdpAdd();
    },
    methods: {
        SetHdp: function (e) {
            this.Status = e
        },
        HdpAdd: function () {    //幻灯片
            this.Status++;
            this.Status = (this.Status >= this.Len) ? 0 : this.Status;
            setTimeout(this.HdpAdd, this.Time);

        },
    },
}
</script>
<style scoped>
.myoo-carousel {
    width: 100%;
    height: 100%;
}

.carousel-content {
    height: 100%;
    width: 100%;
    display: flex;
    position: relative;
    transition: all 0.3s;
}

.carousel-trig {
    position: absolute;
    bottom: 5px;
    right: 10px;
}

.carousel-content li {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
}

.carousel-content img {
    vertical-align: middle;
    height: 100%;
    width: 100%;
    object-fit: cover;
}

.carousel-content .text {
    width: 100%;
    border-radius: 4px;
    background: -webkit-linear-gradient(transparent, rgba(0, 0, 0, 0.5));
    color: #fff;
    font-size: 15px;
    padding: 5px 12px;
    position: absolute;
    text-align: left;
    bottom: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.carousel-trig li {
    cursor: pointer;
    margin: 0 3px;
    background-color: #fff;
    display: inline-block;
    transition: opacity 1s;
    border-radius: 50%;
    width: 10px;
    height: 10px;
}

.carousel-trig li.on,
.carousel-trig li:hover {
    border: 2px solid #ffffff;
    padding: 4px;
    background-color: #30c0ff;
}
</style>