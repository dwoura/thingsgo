/*
 * @Author: chaoxiaoshu-mx leukotrichia@163.com
 * @Date: 2023-02-21 15:34:42
 * @LastEditors: chaoxiaoshu-mx leukotrichia@163.com
 * @LastEditTime: 2023-10-25 09:42:54
 * @FilePath: \ThingsPanel-Backend-Vue\src\api\alarm.js
 * @Description: 告警信息
 */
import axios from "./interceptor/http.js";

/**
 * @description: 告警信息
 * @return {*}
 */
export default {
    /**
     * 告警信息列表
     * @param {*} data 
     * @returns 
     */
    list: (data) => {
        return axios({
            url: "/v1/warning/information/list",
            method: "post",
            data
        })
    },
    
    /**
     * 告警信息处理
     * @param {*} data 
     * @returns 
     */
    process: (data) => {
        return axios({
            url: "/v1/warning/information/edit",
            method: "post",
            data
        })
    },

    /**
     * 批量处理
     * @param {*} data 
     * @returns 
     */
    batchProcess: (data) => {
        return axios({
            url: "/v1/warning/information/batch/processing",
            method: "post",
            data
        })
    },
    
    /**
     * @description: 获取告警信息总数
     * @return {*}
     */
    getAlarmCount: (data) => {
        return axios({
            url: "/v1/warning/information/count",
            method: "post",
            data
        })
    },


}