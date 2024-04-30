<template>
	<view class="page-body">
		<view class="page-section">
			<map id="map1" class="mapstyle" ref="map1" style="width: 100%; height: 100%;" :markers="markers"
				:latitude="latitude" :longitude="longitude" :scale="mapScale" @markertap="markertap">
				<cover-view slot="callout">
					<block v-for="(item, index) in customCalloutMarkerIds" :key="index">
						<cover-view class="customCallout" :marker-id="item">
							<cover-view class="content">
								<cover-view
									style="font-size: 14px;background-color: white;border: 1px #0d6c6c solid;border-radius: 5px;color: #0d6c6c;">
									<cover-view style="margin: 5px;padding: 3px 5px;border-bottom: 1px #0d6c6c solid;">
										设备编号：2534534534</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">定位时间：2023.01.01 09:43</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">停留时间：2023.01.01 09:43</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">定位时速：1.27km/h</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;white-space:pre-wrap;">定位位置：广东省深圳市南山区科技三中路大查查大厦D栋1905
									</cover-view>
									<cover-view style="margin: 5px;padding: 3px 5px;">定位位置：广东省深圳市南山区给给发个给大厦G栋
									</cover-view>
								</cover-view>
							</cover-view>
						</cover-view>
					</block>
				</cover-view>
			</map>
		</view>
	</view>
</template>
<script>
	export default {
		data() {
			return {
				latitude: 30.17489176432292,
				longitude: 120.2113267686632,
				markers: [{
					id: 1,
					latitude: 30.174892900,
					longitude: 120.2113275343,
					iconPath: '../../../static/img/wz.png',
					width: 30,
					height: 30,
					stationName: '',
					distance: 10,
					customCallout: {
						anchorY: 0,
						anchorX: 0,
						display: 'ALWAYS',
					}
				}, {
					id: 2,
					latitude: 30.174894900,
					longitude: 120.2133285343,
					iconPath: '../../../static/img/wz.png',
					width: 30,
					height: 30,
					stationName: '',
					distance: 20,
					customCallout: {
						anchorY: 0,
						anchorX: 0,
						display: 'NONE',
					}
				}, {
					id: 3,
					latitude: 30.172792900,
					longitude: 120.2133285343,
					iconPath: '../../../static/img/wz.png',
					width: 30,
					height: 30,
					stationName: '',
					distance: 30,
					customCallout: {
						anchorY: 0,
						anchorX: 0,
						display: 'NONE',
					},
				}], // 地图上markers列表
				customCalloutMarkerIds: [1, 2, 3],
				mapScale: 16, // 地图默认放大倍数
			}
		},
		methods: {
			markertap(e) {
				const that = this
				let markers = this.markers
				markers.find(function(item, index) {
					if (item.id == e.markerId) {
						that.swiperCurrent = index // 点击marker 实现底部滑到相对应的站点
						item.customCallout.display = 'ALWAYS' // 点击marker 显示站点名
						item.width = 35
						item.height = 35
					} else {
						item.customCallout.display = 'NONE'
						item.width = 25
						item.height = 25
					}
				})
			}
		}
	}
</script>
<style lang="less" scoped>
	.page-body {
		width: 100%;
		height: 100%;
		position: absolute;
		overflow: hidden;
 
		.page-section {
			width: 100%;
			height: 100%;
			position: absolute;
		}
 
		.customCallout {
			width: 75%;
			box-sizing: border-box;
			border-radius: 4rpx;
			display: inline-flex;
			justify-content: center;
			align-items: center;
		}
	}
 
	.mapstyle {
		position: relative;
	}
</style>