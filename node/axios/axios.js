/*axios异步请求库*/
import axios from 'axios';
import qs from 'qs';
var server = document.getElementById("myoo-server") ? document.getElementById("myoo-server").content : "";
axios.interceptors.request.use((config) => {
    if (config.method == "post") {
        config.data = qs.stringify(config.data);
        config.headers['Content-Type'] = 'application/x-www-form-urlencoded';
    }
    config.url = server + "/rest/v1" + config.url
    return config;
}, (error) => {
    return Promise.reject(error);
});
axios.defaults.headers.common['X-MYOO'] = localStorage.getItem("MYOO") ? localStorage.getItem("MYOO") : "";
export default axios;