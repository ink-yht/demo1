<template>
  <div>
    <form @submit.prevent="uploadFile">
      <label for="fileInput">选择文件：</label>
      <input
        type="file"
        id="fileInput"
        ref="fileInput"
        @change="handleFileChange"
      />
      <br />
      <label for="idCardColInput">身份证号码列号：</label>
      <input type="number" v-model="idCardCol" min="1" />
      <br />
      <label for="disabilityNoColInput">残疾证号码列号：</label>
      <input type="number" v-model="disabilityNoCol" min="1" />
      <br />
      <label for="phoneColInput">手机号码列号：</label>
      <input type="number" v-model="phoneCol" min="1" />
      <br />
      <label for="emailColInput">邮箱列号：</label>
      <input type="number" v-model="emailCol" min="1" />
      <br />
      <!-- 添加获取工作表名称的输入框 -->
      <label for="sheetNameInput">工作表名称：</label>
      <input type="text" v-model="sheetName" />
      <br />
      <input type="submit" value="上传并导出 Excel" />
    </form>
  </div>
</template>
<script lang="ts">
import { ref } from "vue";
import * as XLSX from "xlsx";

// 定义接口来表示列号映射数据结构
interface ColumnMapping {
  id_card_col: number;
  disability_no_col: number;
  phone_col: number;
  email_col: number;
}

export default {
  setup() {
    const file = ref<File | null>(null);

    const idCardCol = ref<number | null>(null);
    const disabilityNoCol = ref<number | null>(null);
    const phoneCol = ref<number | null>(null);
    const emailCol = ref<number | null>(null);

    // 定义用于存储工作表名称的变量
    const sheetName = ref<string>("");

    const handleFileChange = (e: Event) => {
      const target = e.target as HTMLInputElement;
      file.value = target.files?.[0] || null;
    };

    const validateInputs = (): boolean => {
      if (!file.value) {
        alert("请选择文件");
        return false;
      }
      if (
        idCardCol.value === null ||
        disabilityNoCol.value === null ||
        phoneCol.value === null ||
        emailCol.value === null
      ) {
        alert("请填写所有列号");
        return false;
      }
      return true;
    };

    const jsonToExcel = (json: any[], filename: string) => {
      // 定义 Excel 表头
      const headers = [
        ["序号", "身份证号", "残疾号", "邮箱", "手机号", "校验消息"],
      ];

      // 转换 JSON 数据为数组格式
      const data = json.map((item) => [
        item.index, // 序号
        item.id_card, // 身份证号
        item.disability_no, // 残疾号
        item.email, // 邮箱
        item.phone, // 手机号
        item.validation_msg, // 校验消息
      ]);

      // 合并表头和数据
      const worksheetData = headers.concat(data);

      // 创建工作表和工作簿
      const worksheet = XLSX.utils.aoa_to_sheet(worksheetData);
      const workbook = XLSX.utils.book_new();
      XLSX.utils.book_append_sheet(workbook, worksheet, "Sheet1");

      // 生成 Excel 文件并触发下载
      XLSX.writeFile(workbook, filename);
    };

    const uploadFile = async () => {
      if (!validateInputs()) return;

      const formData = new FormData();
      formData.append("file", file.value!);

      const columnMapping: ColumnMapping = {
        id_card_col: idCardCol.value!,
        disability_no_col: disabilityNoCol.value!,
        phone_col: phoneCol.value!,
        email_col: emailCol.value!,
      };
      formData.append("columnMapping", JSON.stringify(columnMapping));

      // 将工作表名称添加到formData中，名字为sheet_name
      formData.append("sheet_name", sheetName.value);

      try {
        const response = await fetch("http://localhost:8088/upload", {
          method: "POST",
          body: formData,
        });

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const jsonData = await response.json();
        console.log("Received JSON Data:", jsonData);

        // 将 JSON 数据导出为 Excel
        jsonToExcel(jsonData, "exported_data.xlsx");

        alert("上传成功，Excel 文件已下载！");
        // 添加以下代码实现页面刷新
        location.reload();
      } catch (error) {
        console.error("上传失败:", error);
        alert("上传失败，请检查网络连接或后端服务。");
      }
    };

    return {
      file,
      idCardCol,
      disabilityNoCol,
      phoneCol,
      emailCol,
      sheetName,
      handleFileChange,
      uploadFile,
    };
  },
};
</script>