<template>
    <view class="page-main">
        <liu-easy-map ref="liuEasyMap" :centerLat="'36.05709'" :centerLng="'103.82538'" :scale="14"
            :markerData="markerData" :customCalloutMarkerIds="customCalloutMarkerIds" :polygons="polygons" @clickMarker="markerClick" ></liu-easy-map>
    </view>
</template>

<script>
    export default {
        data() {
            return {
                markerData: [{
                    id: 1,
                    name: '兰州市政府', //标记点展示名字
                    address: '甘肃省兰州市城关区人民政府',
                    latitude: '36.05989', //标记点纬度
                    longitude: '103.83502', //标记点经度
                    markerUrl: '../../static/marker.png', //标记点图标地址
                    iconWidth: 48, //标记点图标宽度
                    iconHeight: 48, //标记点图标高度
                    calloutColor: '#ffffff', //气泡窗口 文本颜色
                    calloutFontSize: 14, //气泡窗口 文本大小
                    calloutBorderRadius: 6, //气泡窗口 边框圆角
                    calloutPadding: 18, //气泡窗口 文本边缘留白
                    calloutBgColor: '#0B6CFF', //气泡窗口 背景颜色
                    calloutDisplay: 'ALWAYS', //气泡窗口 展示类型 默认常显 'ALWAYS' 常显 'BYCLICK' 点击显示
                }, {
                    id: 2,
                    name: '测试地址', //标记点展示名字
                    address: '测试地址详细地址测试地址详细地址测试地址详细地址测试地址详细地址',
                    latitude: "36.05064",
                    longitude: "103.82538"
                }],
				//自定义气泡绑定在标点的索引  待更新
				customCalloutMarkerIds: [1,2],
                //展示区域点位信息
                polygons: [{
                    points: [{
                        latitude: "36.06637",
                        longitude: "103.82471"
                    }, {
                        latitude: "36.06255",
                        longitude: "103.82321"
                    }, {
                        latitude: "36.06234",
                        longitude: "103.81928"
                    }, {
                        latitude: "36.06036",
                        longitude: "103.82175"
                    }, {
                        latitude: "36.05653",
                        longitude: "103.82152"
                    }, {
                        latitude: "36.05953",
                        longitude: "103.82476"
                    }, {
                        latitude: "36.05690",
                        longitude: "103.82785"
                    }, {
                        latitude: "36.06023",
                        longitude: "103.82747"
                    }, {
                        latitude: "36.06243",
                        longitude: "103.83014"
                    }, {
                        latitude: "36.06245",
                        longitude: "103.82616"
                    }], //经纬度数组
                    strokeWidth: 2, //描边的宽度
                    strokeColor: "#FF000060", //描边的颜色
                    fillColor: "#FF000090" //填充颜色
                }]
            }
        },
        methods: {
            //点击标记点
            markerClick(e) {
                console.log('点击标记点信息：', e)
            }
        }
    }
</script>

<style lang="scss">
    .page-main {
        width: 100%;
        height: 100vh;
    }
</style>