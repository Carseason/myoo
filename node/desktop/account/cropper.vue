<template>
    <div id="cropper">
        <div class="cropper-eptiy">
            <img id="image" src />
        </div>
        <div class="cropper-btn">
            <button @click="CroppersBase64()">上传头像</button>
        </div>
    </div>
</template>

<script>
import Cropper from 'cropperjs';
import 'cropperjs/dist/cropper.css';
export default {
    props: ['imgfile'],
    data() {
        return {
            cropper: null,
            Base64Img: "",
        }
    },
    mounted() {
        this.NewCroppers();
        this.cropper.replace(this.imgfile)
    },
    watch: {
        imgfile() {
            this.cropper.replace(this.imgfile)
        },
    },
    methods: {
        NewCroppers() {
            this.cropper = new Cropper(document.getElementById('image'), {
                aspectRatio: 1 / 1,
                viewMode: 1,
                dragMODE: 'move',
                background: true,
                movable: false,           //移动图像
                cropBoxResizable: true, //禁止改变裁剪框大小
                zoomable: false,         //禁止缩放图片
            });
        },
        CroppersBase64() {
            this.Base64Img = this.cropper.getCroppedCanvas({
                width: 150,
                height: 150,
            }).toDataURL('image/jpeg')
            this.Update();
        },
        Update() {
            this.$axios.post('/account/info', {
                type: "avatar",
                avatar: this.Base64Img,
            }).then(res => {
                console.log(res.data)
                this.$store.commit("msg", res.data.message)
            }).catch(error => {
                console.log(error);
            });
            return;
        }
    }
}
</script>
<style  scoped>
#cropper {
    width: 100%;
}
.cropper-eptiy {
    width: 100%;
    height: 380px;
    border-style: dashed;
    border-width: 1px;
    overflow: hidden;
}
.cropper-hidden {
    display: none !important;
}

.cropper-btn {
    text-align: center;
    margin-top: 40px;
}
.cropper-btn button {
    color: #fff;
    width: 110px;
    height: 30px;
    background: #00a1d6;
    border: none;
    border-radius: 4px;
}
</style>
