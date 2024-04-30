import {business_add, business_delete, business_edit} from "@/api/business";
import {message_success} from "@/utils/helpers";
import {dateFormat} from "@/utils/tool";
import {ref, reactive} from "@vue/composition-api";

export default function useBusinessCUD(tableData){
    let loading = ref(false)

    // 创建业务
    function handleCreate(){
        // 创建时添加一个空数据到 tableData
        tableData.value.unshift({
            id: Date.now(),
            name: "",
            created_at: "",
            status: "creating",
            errors: {
                name: ""
            },
            formData: {
                name: ""
            }
        })

    }

    // 修改业务名
    function handleEdit(item){
        // 标记当前项为编辑状态
        item.status = "editing"
        // 表单赋值
        item.formData.name = item.name
    }

    // 取消编辑或新建
    function handleCancel(item){
        if(item.status === "creating") {
            // 取消创建的时候删除本条数据
            let index = tableData.value.indexOf(item)
            tableData.value.splice(index, 1)
        }
        item.status = null
    }

    // 保存编辑或新建
    function handleSave(item){
        // 每次提交前前先清除错误
        item.errors.name = ""

        // 验证
        if(!item.formData.name){
            item.errors.name = "请填写名称"
            return true
        }

        if(loading.value) return
        loading.value = true

        // 提交
        if(item.status === "editing"){
            business_edit({id: item.id, name:item.formData.name}).then(({data})=>{
                if(data.code === 200){
                    message_success("修改成功")
                    // 跟新后 data 返回空，所以用 item.formData
                    item.name = item.formData.name
                    item.status = null
                }
            }).finally(()=>{
                loading.value = false
            })
        } else if(item.status === "creating") {
            business_add({name: item.formData.name}).then(({data})=>{
                if(data.code === 200) {
                    // 洗一下数据 格式化日期
                    item.id = data.data.id
                    item.name = data.data.name
                    item.created_at = dateFormat(data.data.created_at)
                    item.status = null
                    message_success("创建成功")
                }
            }).finally(()=>{
                loading.value = false
            })
        }

    }

    // 删除
    function handleDelete(item){
        if(loading.value) return
        loading.value = true

        business_delete({id: item.id}).then(({data})=>{
            if(data.code === 200) {
                let index = tableData.value.indexOf(item)
                tableData.value.splice(index, 1)
                message_success("删除成功")
            }
        }).finally(()=>{
            loading.value = false
        })
    }

    return {
        handleCreate,
        handleEdit,
        handleCancel,
        handleSave,
        handleDelete,
    }
}