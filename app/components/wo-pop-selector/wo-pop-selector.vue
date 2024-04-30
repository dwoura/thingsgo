<template>
	<view class="comp-layout">
      <view @click="onSelector">
        <slot></slot>
      </view>
      <view class="pop-selector" v-if="state.isShow" :style="state.baseStyle">
				<slot name="header"></slot>
        <scroll-view scroll-y="true" :style="{ height: props.maxHeight }">
          <view
            class="selector-item"
            v-for="item in state.options"
            :key="item.value"
            @click="onChange(item)"
						:style="{ color: state.selectedValue === item.value ? props.activeColor : '' }"
            :class="{ 'selector-item-disabled': item.disabled }">
            <view style="padding-right: 90rpx">{{ item.label }}</view>
						<view v-if="state.selectedValue === item.value">
							<slot name="right">
								✓
							</slot>
						</view>
          </view>
        </scroll-view>
      </view>
    </view>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue';
const props = withDefaults(
  defineProps<{
    defaultValue: string | number;
    selectOptions: any[];
		cardStyle: object;
		activeColor: string;
		maxHeight: string;
  }>(),
  {
    defaultValue: 0,
    selectOptions: () => [
      {
        value: 0,
        label: '选项1',
      },
      {
        value: 1,
        label: '选项2',
      },
      {
        value: 2,
				disabled: true,
        label: '选项3',
      },
      {
        value: 3,
        label: '选项4',
      }
    ],
		cardStyle: () => {
			return {
				background: '#ffffff',
				border: '1px solid #ebeef5',
				borderRadius: '6px',
				boxShadow: '0 2px 12px 0 rgba(0, 0, 0, 0.1)',
				padding: '4px 0',
				fontSize: '26rpx'
			}
		},
		activeColor: '#58BA86',
		maxHeight: '250rpx'
  }
);
watch(
  () => props.cardStyle,
  (newVal, oldVal) => {
    state.baseStyle = newVal;
  }
);
const $emit = defineEmits<{
  (e: 'changeSelect', ev: any): void;
}>();

const state = reactive({
  isShow: false,
  options: props.selectOptions,
  selectedValue: props.defaultValue,
	baseStyle: props.cardStyle,
});

const onSelector = () => {
  state.isShow = !state.isShow;
};

const onChange = (data: any) => {
  if (data.disabled) {
    return;
  }
  state.selectedValue = data.value;
  state.isShow = false;
  $emit('changeSelect', data);
};
</script>

<style scoped>
.comp-layout {
  position: relative;
}
.pop-selector {
  position: absolute;
  top: calc(100% + 4px);
  right: 0;
  z-index: 2;
}
.selector-item {
  display: flex;
  justify-content: space-between;
  text-align: center;
  padding: 15rpx 20rpx;
  width: auto;
  white-space: nowrap;
  flex-wrap: nowrap;
}
.selector-item-disabled {
  opacity: 0.4;
  cursor: default;
}
</style>
