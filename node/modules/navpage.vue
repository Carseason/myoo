<template>
    <nav v-if=" maxpage > 1 " class="pagination">
        <template v-if="status == 1">
            <template v-if=" Pages > 1">
                <a class="next page-numbers" @click="--Pages" title="上一页">{{Left}}</a>
                <a @click="Pages = 1">1</a>
            </template>
            <select class="page" v-model.number="Pages">
                <template v-for="item in maxpage">
                    <option :key="item" :value="item">{{item}}</option>
                </template>
            </select>
            <template v-if=" Pages < maxpage">
                <a @click="Pages =maxpage ">{{maxpage}}</a>
                <a class="next page-numbers" title="下一页" @click=" ++Pages ">{{Right}}</a>
            </template>
        </template>

        <template v-else>
            <a @click="--Pages" v-if=" Pages > 1" title="上一页">{{Left}}</a>
            <template v-if="maxpage < 10">
                <a v-for="value in maxpage" @click="Pages = value" :key="value" :class="{ on : Pages == value }">{{value}}</a>
            </template>
            <template v-else>
                <template v-if=" Pages < 7 ">
                    <a v-for="value in 10" @click="Pages = value" :key="value" :class="{ on : Pages == value }">{{value}}</a>
                    <a @click="Pages = Pages + 5" class="pagination-rights"></a>
                    <a @click="Pages = maxpage" :class="{ on : Pages == maxpage }">{{maxpage}}</a>
                </template>
                <template v-else-if="( Pages + 10 ) < maxpage">
                    <a @click="Pages = 1" :class="{ on : Pages == 1 }" v-if=" Pages != 1">1</a>
                    <a @click="Pages = Pages - 5 " class="pagination-lefts"></a>

                    <a @click="Pages = (Pages - 5) " :class="{ on : Pages == (Pages - 5)  }">{{(Pages - 5) }}</a>
                    <a @click="Pages = (Pages - 4) " :class="{ on : Pages == (Pages - 4)  }">{{(Pages - 4) }}</a>
                    <a @click="Pages = (Pages - 3) " :class="{ on : Pages == (Pages - 3)  }">{{(Pages - 3) }}</a>
                    <a @click="Pages = (Pages - 2) " :class="{ on : Pages == (Pages - 2)  }">{{(Pages - 2) }}</a>
                    <a @click="Pages = (Pages - 1) " :class="{ on : Pages == (Pages - 1)  }">{{(Pages - 1) }}</a>

                    <a @click="Pages = Pages " :class="{ on : Pages == Pages  }">{{ Pages }}</a>
                    <a v-for="value in 5" :key="(Pages + value )" @click="Pages = (Pages + value ) " :class="{ on : Pages == (Pages + value )  }">{{ (Pages + value ) }}</a>

                    <a @click="Pages = Pages + 5" class="pagination-rights"></a>
                    <a @click="Pages = maxpage" :class="{ on : Pages == maxpage }" v-if=" Pages != maxpage">{{maxpage}}</a>
                </template>
                <template v-else>
                    <a @click="Pages = 1" :class="{ on : Pages == 1 }" v-if=" Pages != 1">1</a>
                    <a @click="Pages = Pages -5" class="pagination-lefts"></a>

                    <a @click="Pages = (Pages - 5) " :class="{ on : Pages == (Pages - 5)  }">{{(Pages - 5) }}</a>
                    <a @click="Pages = (Pages - 4) " :class="{ on : Pages == (Pages - 4)  }">{{(Pages - 4) }}</a>
                    <a @click="Pages = (Pages - 3) " :class="{ on : Pages == (Pages - 3)  }">{{(Pages - 3) }}</a>
                    <a @click="Pages = (Pages - 2) " :class="{ on : Pages == (Pages - 2)  }">{{(Pages - 2) }}</a>
                    <a @click="Pages = (Pages - 1) " :class="{ on : Pages == (Pages - 1)  }">{{(Pages - 1) }}</a>
                    <a @click="Pages = Pages " :class="{ on : Pages == Pages        }">{{ Pages }}</a>
                    <a v-for="value in (maxpage - Pages)" @click="Pages = (value + Pages)" :key="(value + Pages)" :class="{ on : Pages == (value + Pages) }">{{ value + Pages }}</a>
                </template>
            </template>
            <a @click="++Pages" v-if=" Pages < maxpage" title="下一页">{{Right}}</a>
        </template>
    </nav>
</template>

<script>
export default {
    name: 'page',
    props: ['maxpage', 'page', 'status'],
    data: function () {
        return {
            Pages: this.page,
            PageGap: this.maxpage - this.page,
            Left: "<",
            Right: ">",
        }
    },
    watch: {
        page: function (val) {
            if (this.Pages <= 0 || this.Pages > this.maxpage) {
                this.Pages = 1
                this.$emit('update:page', 1)
                return
            }
            this.Pages = val;
            return
        },
        Pages: function (val) {
            this.$emit('update:page', val)
        }
    },
    methods: {

    }
}
</script>
<style scoped>
nav.pagination {
    display: block;
    width: 100%;
    text-align: center;
    margin: 0 auto;
    padding: 10px;
}
nav.pagination a,
nav select.page {
    display: inline-block;
    line-height: 40px;
    height: 40px;
    padding: 0 15px;
    text-align: center;
    border: none;
    margin: 0 5px;
    background-color: #f4f4f5;
    color: #606266;
    font-size: 13px;
    border-radius: 2px;
}
nav.pagination a {
    min-width: 30px;
}
nav select.page {
    min-width: 50px;
}
nav.pagination a.on {
    background-color: #409eff;
    color: #fff;
}
nav.pagination a:hover {
    color: #409eff;
}
nav.pagination a.on:hover {
    color: #fff;
}
.pagination-lefts:before,
.pagination-rights:before {
    content: "...";
}
.pagination-lefts:hover:before {
    content: "◀";
}
.pagination-rights:hover:before {
    content: "▶";
}
</style>

