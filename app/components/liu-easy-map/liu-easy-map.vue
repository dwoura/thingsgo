<template>
	<view style="width: 100%; height: 100%;">
		
		<map style="width: 100%; height: 100%;" id="esaymap" :scale="scale" :latitude="nowLat ? nowLat : centerLat"
			:longitude="nowLng ? nowLng : centerLng" :markers="markers" :polygons="polygonsData"
			:enable-zoom="isEnableZoom" :enable-scroll="isEnableScroll" :enable-satellite="isShowWxMap"
			:enable-rotate="isEnableRotate" @markertap="chooseItem" @tap="clickMap">
			
			<view style="display: flex;justify-content: center;padding:10px;">
				<view style="display: flex;" >
					<epselect :disabled="false" v-model="selectBusiness" :options="optionsBusiness"  @change="selectOptionsBusiness"></epselect>
				</view>
				<view style="display: flex;" >
					<epselect :disabled="false" v-model="selectDevice" :options="optionsDevice"  @change="selectOptionsDevice"></epselect>
				</view>
			</view>
			

			
			<cover-view slot="callout">
					<block v-for="(item, index) in customCalloutMarkerIds" :key="index">
						<cover-view class="customCallout" :marker-id="item">
							<cover-view class="content">
								<cover-view v-if="isShowTHCallout"
									style="font-size: 14px;background-color: white;border: 1px #0d6c6c solid;border-radius: 5px;color: #0d6c6c;">
									<cover-view style="margin: 5px;padding: 3px 5px;border-bottom: 1px #0d6c6c solid;">
										项目名：{{calloutData.groupName || "项目名"}} —— 设备名：{{calloutData.deviceName || "设备名"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">soil：{{calloutData.thData.soil || "36"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">光照：{{calloutData.thData.lux || "8.2Klux"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">土壤ph值：{{calloutData.thData.soilPh || "6.91"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">土壤温度：{{calloutData.thData.soilT || "1.9℃"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">土壤湿度：{{calloutData.thData.soilH || "7.38%rh"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">温度：{{calloutData.thData.t || "-6.09℃"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">湿度：{{calloutData.thData.h || "7.38"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">风向：{{calloutData.thData.wd || "7.38"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">风速：{{calloutData.thData.ws || "7.38"}}</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">数据更新时间：{{calloutData.thData.systime || "2024-04-22 13:02:56"}}  </cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;white-space:pre-wrap;">定位位置：xxx
									</cover-view>
								</cover-view>
							</cover-view>
						</cover-view>
					</block>
			</cover-view>
			
		</map>
		
		<view class="rightbox">
<!-- 			<view class="boxitem" @click="changeTab(1)">
				<image class="itemimg" :src="tabIndex ? myaddressOnImg : myaddressImg" mode=""></image>
				<view class="itemname" :class="tabIndex ? 'active' : ''">我的位置</view>
			</view> -->
			<view class="boxitem" @click="changeTab(2)" v-if="wxMapShow">
				<image class="itemimg" :src="tabIndex2 ? wxmapOnImg:wxmapImg" mode=""></image>
				<view class="itemname" :class="tabIndex2 ? 'active' : ''">卫星地图</view>
			</view>
		</view>
		<cover-view class="detailbox" v-if="isShowDetail">
			<cover-image class="closeicon" :src="closeImg" @click="closeDetail"></cover-image>
			<cover-view class="boxl">
				<cover-view class="boxlhd">{{detailData.name || '--'}}</cover-view>
				<cover-view class="boxlbd">{{detailData.address || '--'}}</cover-view>
			</cover-view>
			<cover-view class="boxr" @click="toDeviceDetail(detailData.item)">
				<cover-view class="boxrhd">{{'设备详情'}}</cover-view>
				<cover-image class="boxrimg" :src="goImg" mode=""></cover-image>
			</cover-view>
		</cover-view>
	</view>
</template>

<script>
import epselect from 'vuex';
	export default {


		props: {
			//中心点纬度
			centerLat: {
				type: String,
				default: ''
			},
			//中心点经度
			centerLng: {
				type: String,
				default: ''
			},
			//标记点数据
			markerData: {
				type: Array,
				default () {
					return []
				}
			},
			//多边形数据
			polygons: {
				type: Array,
				default () {
					return []
				}
			},
			//标记点图标宽度
			markerIconWidth: {
				type: Number,
				default: 48
			},
			//标记点图标高度
			markerIconHeight: {
				type: Number,
				default: 48
			},
			//标记点图标路径
			markerIconUrl: {
				type: String,
				default: ''
			},
			//缩放级别 取值范围为3-20
			scale: {
				type: Number,
				default: 16
			},
			//是否显示指南针
			isShowCompass: {
				type: Boolean,
				default: false
			},
			//是否支持缩放
			isEnableZoom: {
				type: Boolean,
				default: true
			},
			//是否支持拖动
			isEnableScroll: {
				type: Boolean,
				default: true
			},
			//是否支持旋转
			isEnableRotate: {
				type: Boolean,
				default: false
			},
			//marker索引
			// customCalloutMarkerIds: {
			// 	type: Array,
			// 	default () {
			// 		return []
			// 	}
			// },
		},
		watch: {
			markerData: {
				immediate: true, //初始化的时候是否调用
				deep: true, //是否开启深度监听
				handler(newValue, oldValue) {
					this.markerDatas = newValue
					this.showMarkers()
				}
			},
			polygons: {
				immediate: true, //初始化的时候是否调用
				deep: true, //是否开启深度监听
				handler(newValue, oldValue) {
					this.polygonsData = [...newValue]
				}
			}
		},
		data() {
			return {
				markerImg: require('../../static/marker.png'),
				goImg: require('../../static/go.png'),
				myaddressImg: require('../../static/myaddress.png'),
				wxmapImg: require('../../static/wxmap.png'),
				myaddressOnImg: require('../../static/myaddress-on.png'),
				wxmapOnImg: require('../../static/wxmap-on.png'),
				closeImg: require('../../static/close.png'),
				polygonsData: [], //polygons区域数据
				markers: [], //markers数据
				customCalloutMarkerIds: [],
				detailData: {}, //选中展示详情数据
				nowLat: '', //我的当前位置
				nowLng: '',
				tabIndex: false,
				tabIndex2: false,
				isShowWxMap: false, //是否展示卫星地图 
				isShowDetail: false, //是否展示详情弹框
				wxMapShow: false, //是否展示卫星地图按钮（小程序展示）
				activeMarkerId: null,  // 存储当前显示气泡的标记点 ID
				ywData: [], // 下拉列表业务数据
				//选择设备分组 下拉列表
				selectBusiness:"0",
				optionsBusiness: [{
					value: "0",
					label: "请选择设备分组"
				}],
				//选择设备 下拉列表
				selectDevice:"0",
				optionsDevice: [{
					value: "0",
					label: "再选择设备"
				}],
				deviceList:[],
				currentGroup:"",
				currentDevice:"",
				assetData: null,
				//温湿度传感器气泡展示
				isShowTHCallout: true, 
				//抓拍气泡展示
				isShowSnapCallout: true, 
				//开关气泡展示
				isShowSWCallout: true, 
			}
		},
		mounted() {
			const type = uni.getSystemInfoSync().uniPlatform
			if (type == 'mp-weixin') {
				this.wxMapShow = true
			}
			this.showMarkers()
			if (!this.centerLat || !this.centerLng) this.getLocation()
			
			//展示业务集合
			this.getYwData()
			//获取所有设备的列表 用于后续根据业务筛选
			//this.getDeviceList();
		},
		methods: {
			// 获取业务列表
			getYwData() {
				uni.showLoading({
					title: '加载中'
				});
				this.API.apiRequest('/api/business/index', {
					page: 1,
					limit: 10
				}, 'post').then(res => {
					if (res.code === 200) {
						this.ywData = res.data.data //获取业务数据
						console.log("yw",this.ywData)
						//为每个业务索引获取业务详情
						let tempList = [] 
							this.ywData.forEach((item,index)=>{
									//请求api获取每个组对应的asset_id
									this.API.apiRequest('/api/asset/list/d', {
										business_id: item.id
									}, 'post').then(res => {
										if (res.code === 200) {
											if (res.data && res.data.length > 0) {
												item.secondShow = !item.secondShow
												const data = res.data
												data.forEach(t => {
													t.device_group = t.device_group.replace(/\//g, '');
												})
												this.ywData.forEach(d => {
													if (item.id == d.id) {
														d.equipLists = data
													}
												})
												
												// console.log("aaa", this.currentGroup);
												// if (!this.currentGroup) {
												// 	console.log("currentGroup")
												// 	this.currentGroup = data[0]
												// 	this.deviceList = []
												// 	this.getDeviceList()
												// }
												//this.$forceUpdate()
												
												
												tempList.push({
													value: index,
													label: item.name,
													type: 'business',
													assetId: data[0].id,
													//item: item
												})
												
												//this.$forceUpdate()
											}
										}
										// setTimeout(()=>{
										// uni.hideLoading();
										// },500);
									});
									
									
							});
						this.optionsBusiness = tempList
						this.selectBusiness = "请选择设备分组"
					}
				}).finally(() => {
					// uni.hideLoading()
				});
				setTimeout(() => {
					uni.hideLoading()
				}, 300);
			},
			// 获取设备列表
			getDeviceList() {
				// 清除定时器
				clearInterval(this.timer)
				this.markerData = [{}] //重置标点
				var newData = {};
				console.log("获取设备列表里",this.currentGroup.assetId)
				this.API.apiRequest('/api/device/list', {
					asset_id: this.currentGroup.assetId,
					current_page: this.$store.state.list.equpPage,
					per_page: 20
				}, 'post').then(res => {
					if (res.code === 200) {
						var newData = res.data.data || [];
						var data = []
						if (newData.length > 0) {
							newData.forEach(item => {
								// if (item.device_type != 2) {
								// 	data.push(item)
								// }
								data.push(item)
							})
						}
						this.deviceList = data;
						//更新二级下拉列表 更新标点信息
						
						this.selectDevice = "0"
						let tempCustomCalloutMarkerIds = [] //气泡索引重置
						let tempList = []
						let tempMarkerList = []
						this.deviceList.forEach((item,index)=>{
							tempList.push({
								value: index,
								label: item.device_name,
								item: item
							})
							tempCustomCalloutMarkerIds.push(index) //气泡索引逐步添加
							//更新标点信息
							if (item.location!=""){
								let locations = item.location.split(",") 
								//console.log("l",locations)
								tempMarkerList.push ({
									id: index,
									name: item.device_name, //标记点展示名字
									address: '暂无文本',
									longitude: locations[0],
									latitude: locations[1],
									markerUrl: '../../static/marker.png', //标记点图标地址
									customCallout: {
										anchorY: 0,
										anchorX: 0,
										display: 'BYCLICK',
									},
									item: item
								})
							}
						});
						
						this.optionsDevice = tempList
						this.nowLng = tempMarkerList[0]?.longitude
						this.nowLat = tempMarkerList[0]?.latitude
						this.markerData = tempMarkerList
						this.customCalloutMarkerIds = tempCustomCalloutMarkerIds
						//console.log(this.deviceList)
						
						//加载气泡索引与信息
						
						//this.showMarkers();
						this.$forceUpdate();//刷新
					}
				})
			},
			//右侧类型切换
			changeTab(index) {
				if (index == 1) {
					this.tabIndex = !this.tabIndex
					if (this.tabIndex) this.getLocation()
					else this.showMarkers()
				} else {
					this.tabIndex2 = !this.tabIndex2
					if (this.tabIndex2) this.isShowWxMap = true
					else this.isShowWxMap = false
				}
			},
			//获取当前的地理位置
			getLocation() {
				uni.getLocation({
					type: 'gcj02',
					isHighAccuracy: true,
					highAccuracyExpireTime: 3500,
					success: (res) => {
						this.nowLat = res.latitude
						this.nowLng = res.longitude
						let arr = [{
							id: 9999,
							latitude: res.latitude || '', //纬度
							longitude: res.longitude || '', //经度
							width: this.markerIconWidth, //宽
							height: this.markerIconHeight, //高
							iconPath: this.markerImg
						}];
						this.markers = [...arr];
						let mapObjs = uni.createMapContext('esaymap', this)
						uni.setStorage({
							mapObjs: mapObjs
						})
						mapObjs.moveToLocation({
							latitude: res.latitude,
							longitude: res.longitude
						}, {
							complete: res => {}
						})
					},
					fail: (res) => {
						if (res.errMsg == "getLocation:fail auth deny") {
							uni.showModal({
								content: '检测到您没打开获取信息功能权限，是否去设置打开？',
								confirmText: "确认",
								cancelText: '取消',
								success: (res) => {
									if (res.confirm) {
										uni.openSetting({
											success: (res) => {}
										})
									} else {
										return false;
									}
								}
							})
						}
					}
				})
			},
			//到这去
			goRoute() {
				uni.openLocation({
					latitude: +this.detailData.latitude,
					longitude: +this.detailData.longitude,
					scale: 17,
					name: this.detailData.name || '--',
					address: this.detailData.address || '--'
				});
			},
			//地图打点展示marker
			showMarkers() {
				if (this.markerDatas && this.markerDatas.length > 0) {
					var arr = []
					for (var i = 0; i < this.markerDatas.length; i++) {
						arr.push({
							id: Number(this.markerDatas[i].id),
							latitude: this.markerDatas[i].latitude || '', //纬度
							longitude: this.markerDatas[i].longitude || '', //经度
							iconPath: this.markerDatas[i].markerUrl ? this.markerDatas[i].markerUrl : this
								.markerImg, //显示的图标        
							rotate: 0, // 旋转度数
							width: this.markerDatas[i].iconWidth ? this.markerDatas[i].iconWidth : this
								.markerIconWidth, //宽
							height: this.markerDatas[i].iconHeight ? this.markerDatas[i].iconHeight : this
								.markerIconHeight, //高
							// callout: { //自定义标记点上方的气泡窗口 点击有效
							// 	content: this.markerDatas[i].name, //文本
							// 	color: this.markerDatas[i].calloutColor || '#ffffff', //文字颜色
							// 	fontSize: this.markerDatas[i].calloutFontSize || 14, //文本大小
							// 	borderRadius: this.markerDatas[i].calloutBorderRadius || 6, //边框圆角
							// 	padding: this.markerDatas[i].calloutPadding || 6,
							// 	bgColor: this.markerDatas[i].calloutBgColor || '#0B6CFF', //背景颜色
							// 	display: this.markerDatas[i].calloutDisplay || 'BYCLICK', //常显
							// },
							
							customCallout: {
								anchorY: 0,
								anchorX: 0,
								display: 'BYCLICK',
							}
						})
					}
					this.markers = arr
				}
			},
			//点击标记点
			chooseItem(e) {
				let markerId = e.detail.markerId
				for (var i = 0; i < this.markerDatas.length; i++) {
					if (this.markerDatas[i].id == markerId) {
						this.isShowDetail = true
						this.detailData = this.markerDatas[i]
						this.$emit("clickMarker", this.markerDatas[i])
						break
					}
				}
				
				const that = this
				let markers = this.markers
				markers.find(function(item, index) {
					if (item.id == e.markerId && item.customCallout.display != 'ALWAYS') {
						that.swiperCurrent = index // 点击marker 实现底部滑到相对应的站点
						item.customCallout.display = 'ALWAYS' // 点击marker 显示站点名
						item.width = 64
						item.height = 64
					}else if (item.id == e.markerId && item.customCallout.display == 'ALWAYS'){
						item.customCallout.display = 'NONE'
						item.width = 48
						item.height = 48
					}
				})
			},
			//点击地图(仅微信小程序支持)
			clickMap(e) {
				// #ifdef MP-WEIXIN
				let lat = e.detail.latitude.toFixed(5)
				let lng = e.detail.longitude.toFixed(5)
				this.$emit("clickMap", {
					latitude: lat,
					longitude: lng
				})
				//console.log("点击了地图")
				// #endif
			},
			//关闭详情弹框
			closeDetail() {
				this.detailData = {}
				this.isShowDetail = false
			},
			//选择列表框后展示设备分组list
			showBusinessList(){

			},
			//选中设备分组选项
			selectOptionsBusiness(e){
				//2 展示该分组下所有设备坐标
				//3 更新设备下拉列表
				//console.log("选中后的选项",e)
				this.selectBusiness = e.value
				// 通过设置currentGroup再获取列表
				this.currentGroup = e
				//console.log("当前组",this.currentGroup)
				//this.deviceList = [] //先清空设备列表
				this.getDeviceList()
				
				console.log("设备列表",this.deviceList)//待解决
				
			},
			//选择列表框后展示设备list
			showDeviceList(){
				
			},
			//选中设备选项
			selectOptionsDevice(e){
				//定位到该设备所在位置
				//展示该设备气泡
				
				this.selectDevice = e.value
				console.log("e",e)

				//定位到该设备位置
				if (e.item.location!=""){
					let locations = e.item.location.split(",")
					let longitude = locations[0]
					let latitude = locations[1]
					this.nowLng = longitude
					this.nowLat = latitude
					// let mapObjs = uni.getStorage('mapObjs')
					// mapObjs.moveToLocation({
					// 	latitude: latitude,
					// 	longitude:longitude
					// }, {
					// 	complete: res => {}
					// })
				}

			},
			showCustomCallout(){
				//判断不同设备，展示不同数据
				//抓拍设备提供跳转抓拍页面，温湿度则展示实时数据、还有开关设备都有详情页面
			},
			// 跳转设备详情页
			toDeviceDetail(data) {
				//判断设备活跃情况
				// var state = ''
				// if (data.latest_ts && this.TimeDifference(this.formatDate(data.latest_ts), this.formatDate(parseInt(
				// 	new Date().getTime() *
				// 	1000))) > 30) {
				// 	state = 0
				// }
				// if (data.latest_ts && this.TimeDifference(this.formatDate(data.latest_ts), this.formatDate(parseInt(
				// 	new Date().getTime() *
				// 	1000))) <= 30) {
				// 	state = 1
				// }
				//判断是否是抓拍设备后再跳转对应页面
				console.log("data",data)
				
				uni.switchTab({
				    url: '/pages/fishery-monitor/fishery-monitor',
				    success: function () {
				        // 在switchTab的success回调中，再使用navigateTo打开子包内的页面
						uni.navigateTo({
							url: '../sub/fishery-monitor/deviceDetail?type=' + data.type + '&device_id=' + data.device_id + '&device_name=' +
								data.device_name + '&latest_ts_name=' + data.latest_ts_name + '&state=' + parseInt(data.status)
						})
				    }
				});

			},
		}
	}
</script>

<style>
	.rightbox {
		padding: 0 8rpx;
		background: #FFFFFF;
		box-shadow: 0rpx 4rpx 8rpx 0rpx rgba(200, 200, 200, 0.5);
		border-radius: 14rpx;
		position: fixed;
		top: 154rpx;
		right: 20rpx;
	}

	.boxitem {
		display: flex;
		flex-direction: column;
		text-align: center;
		padding-bottom: 8rpx;
		border-bottom: 2rpx solid #E4E4E4;
	}

	.itemimg {
		width: 40rpx;
		height: 40rpx;
		margin: 16rpx auto 4rpx;
	}

	.itemname {
		font-size: 22rpx;
		font-weight: 400;
		color: #333333;
		line-height: 42rpx;
	}

	.active {
		color: #2765F1;
	}

	.detailbox {
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: calc(100% - 128rpx);
		padding: 24rpx 32rpx;
		background: #FFFFFF;
		border-radius: 16rpx;
		position: fixed;
		bottom: 32rpx;
		left: 32rpx;
	}

	.closeicon {
		width: 40rpx;
		height: 40rpx;
		position: absolute;
		right: 16rpx;
		top: 12rpx;
	}

	.boxl {
		width: calc(100% - 84rpx);
	}

	.boxlhd {
		margin-bottom: 16rpx;
		white-space: pre-wrap;
		font-size: 36rpx;
		font-weight: bold;
		color: #333333;
		line-height: 48rpx;
	}

	.boxlbd {
		font-size: 30rpx;
		font-weight: 400;
		color: #333333;
		line-height: 46rpx;
		white-space: pre-wrap;
	}

	.boxr {
		width: 350rpx;
		display: flex;
		align-items: center;
		position: relative;
	}

	.boxr::before {
		width: 2rpx;
		height: 96rpx;
		background: #e3e3e3;
		content: "";
		position: relative;
		left: 0;
		z-index: 99;
	}
	
	.boxrhd{
		margin-bottom: 12rpx;
		white-space: pre-wrap;
		font-size: 36rpx;
		font-weight: bold;
		color: #0B6CFF;
		line-height: 48rpx;
		margin: 2rpx;
	}
	
	.boxrimg {
		width: 64rpx;
		height: 64rpx;
		margin: 0 auto;
	}
	
	.customCallout {
		width: 100%;
		box-sizing: border-box;
		border-radius: 4rpx;
		display: inline-flex;
		justify-content: center;
		align-items: center;
	}

</style>