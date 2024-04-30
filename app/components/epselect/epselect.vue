<template>
	<view class="ep-picker-box">
		<!-- 蒙版区域 开始 -->
		<view class="ep-mask-box" v-show="show" @click="show=false"></view>
		<!-- 蒙版区域 开始 -->
 
		<!-- 输入框区域 开始 -->
		<view class="ep-input-box" @click="openOptions" :class="{'disabled':disabled}">
			<view style="display: flex;align-items: center;min-height: 36px;">{{showLabel}}</view>
			<text v-if="!show" class="iconfont icon-xiala"></text>
			<text v-else class="iconfont icon-xialashang"></text>
		</view>
		<!-- 输入框区域 结束 -->
 
		<!-- 弹出的下拉区域 开始 -->
		<view v-show="show" class="ep-picker-content-wrap">
			<scroll-view class="ep-picker-content" :scroll-y="true">
				<!-- 展示下拉项列表 开始 -->
				<view v-for="item in options" :key="item[value_key]"
					:class="{'disabled':item.disabled,'active':value==item[value_key]}" class="option-item"
					@click="itemClick(item)">{{item[label_key]}}</view>
				<!-- 展示下拉项列表 结束 -->
 
				<!-- 下拉列表为空 开始 -->
				<view v-if="options.length==0" class="option-no-data">无数据</view>
				<!-- 下拉列表为空 结束 -->
			</scroll-view>
			<text class="triangle"></text>
		</view>
		<!-- 弹出的下拉区域 结束 -->
	</view>
</template>
 
<script>
	export default {
		data() {
			return {
				show: false,
				left: 0,
 
			}
		},
		props: {
			value: {
				type: [String, Number],
				default: ""
			},
			options: {
				type: Array,
				default: function() {
					return []
				}
			},
			value_key: {
				type: String,
				default: "value"
			},
			label_key: {
				type: String,
				default: "label"
			},
			disabled: {
				type: Boolean,
				default: false
			}
		},
		model: {
			prop: 'value',
			event: "valChange"
		},
		methods: {
			//点击选中选项
			itemClick(item) {
				if (item.disabled) return
 
				//关闭
				console.log("选中后关闭下拉框");
				this.show = false
				console.log(item.value);
				//修改v-model的值
				this.$emit('valChange', item.value)
				//将事件通知父组件
				this.$emit('change', item)
			},
			//展开选项
			openOptions() {
				console.log("展开选项");
				if (!this.disabled) {
					this.show = true
				}
			}
		},
		computed: {
			showLabel() {
				var index = this.options.findIndex(item => {
					return item[this.value_key] == this.value
				})
				if (index != -1) {
					return this.options[index][this.label_key]
				} else if (!this.value) {
					return "全部"
				} else {
					return this.value
				}
			}
		}
	}
</script>
 
<style scoped>
	/* 引入字体图标 */
	/* @import './iconfont.css'; */
 
	.ep-picker-box {
		/* width:100%; */
		/* box-sizing: border-box; */
		margin-top: 30%;
		position: relative;
		display: flex;
		align-items: center;
		font-size: 24rpx;
		color: #619f9f;
		width: 240rpx;
		height: 48rpx;
		border-radius: 24rpx 24rpx 24rpx 24rpx;
		border: 3rpx solid #619f9f;
		background-color: #FFFFFF;
		
	}
 
	.ep-mask-box {
		position: fixed;
		z-index: 999;
		top: 0;
		right: 0;
		left: 0;
		bottom: 0;
		background: none;
	}
 
	.ep-input-box {
		align-items: center;
		border-radius: 4px;
		padding-left: 10px;
		position: relative;
		cursor: pointer;
		width: 200rpx;
	}
 
	/* 整个下拉组件禁用样式 */
	.ep-input-box.disabled {
		cursor: not-allowed;
		background-color: yellow;
		color: #999;
		border-bottom: 1px solid #000;
	}
 
	/* 展开收起箭头样式 */
	.ep-input-box .iconfont {
		position: absolute;
		top: 50%;
		right: 5px;
		font-size: 24px;
		transform: translateY(-50%);
		color: #B8B8B8;
	}
 
	/* 下拉容器样式 外层 */
	.ep-picker-content-wrap {
 
		width: 100%;
		position: absolute;
		top: 55rpx;
		left: 0;
		z-index: 9999;
		padding-top: 6px;
	}
 
	/* 下拉容器样式 内层 */
	.ep-picker-content-wrap .ep-picker-content {
		background-color: #fff;
		color: #999;
		padding: 3px 0;
		box-shadow: 0 0 20px 5px rgb(0 0 0 / 30%);
		border-radius: 5px;
		max-height: 181px;
 
	}
 
	/* 下拉项通用样式 */
	.ep-picker-content-wrap .ep-picker-content .option-item {
		padding: 16rpx 8rpx 10rpx;
 
		font-size: 24rpx;
		cursor: pointer;
		border-bottom: 1rpx solid #DFDDDD;
	}
 
	.ep-picker-content-wrap .ep-picker-content .option-item:last-child {
		border-bottom: none;
	}
 
	/* 无下拉项数据时样式 */
	.ep-picker-content-wrap .ep-picker-content .option-no-data {
		padding: 8px 18px;
		cursor: text;
		color: #999;
		text-align: center;
	}
 
	/* 鼠标移入下拉项样式 */
	.ep-picker-content-wrap .ep-picker-content .option-item:hover {
		background-color: #f5f7fa;
	}
 
	/* 已选中的下拉项样式 */
	.ep-picker-content-wrap .ep-picker-content .option-item.active {
		/* color: #007AFF; */
	}
 
	/* 禁用的下拉项样式 */
	.ep-picker-content-wrap .ep-picker-content .option-item.disabled {
		color: #c0c4cc;
		border-bottom: 1px solid #000;
	}
 
	.ep-picker-content-wrap .ep-picker-content .option-item.disabled:hover {
		cursor: not-allowed;
	}
 
	/* 下拉容器指示箭头样式 */
	.ep-picker-content-wrap .triangle {
		width: 0;
		height: 0;
		border-top: 6px solid rgba(0, 0, 0, 0);
		border-right: 6px solid rgba(0, 0, 0, 0);
		border-bottom: 6px solid #fff;
		border-left: 6px solid rgba(0, 0, 0, 0);
		position: absolute;
		top: -6px;
		left: 50%;
		transform: translateX(-50%);
		box-sizing: content-box;
	}
</style>