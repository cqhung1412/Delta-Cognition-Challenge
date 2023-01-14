import { useState } from "react";
import { Form, Input, Modal, DatePicker, Upload, message } from "antd";
import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";
import type { RcFile } from "antd/es/upload/interface";

const { Item: FormItem } = Form;

const validateImageType = (file: RcFile) => {
  const isJpg = ["image/jpeg", "image/jpg"].includes(
    file.type
  );
  if (!isJpg) {
    message.error("You can only upload JPG/PNG file!");
  }
  const isLt5M = file.size / 1024 / 1024 < 5;
  if (!isLt5M) {
    message.error("Image must smaller than 5MB!");
  }
  
  return isJpg && isLt5M;
};

const getBase64 = (img: RcFile, callback: (bytes: any) => void) => {
  const reader = new FileReader();
  reader.addEventListener("load", () =>
    callback(new Uint8Array(reader.result as ArrayBuffer))
  );
  reader.readAsArrayBuffer(img);
};

const DogForm: React.FC<any> = ({
  onNewDogPost,
  onFinishFailed,
  onCancel,
  visible,
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [imageUrl, setImageUrl] = useState<string>();
  const [fileList, setFileList] = useState<RcFile[]>([]);

  const beforeUpload = (file: RcFile) => {
    if (validateImageType(file)) {
      setFileList([file]);
    }
    return false;
  };

  const onFinish = (values: any) => {
    // . setLoading(true);
    getBase64(fileList[0], (imageBytes) => {
      values.imageBytes = imageBytes;
      onNewDogPost(values, () => {
        setLoading(false);
      });
    });
  };

  const imageUploadButton = (
    <div>
      {loading ? <LoadingOutlined /> : <PlusOutlined />}
      <div style={{ marginTop: 8 }}>Upload</div>
    </div>
  );

  const imageFormItem = (
    <FormItem
      label="Image"
      name="image"
      rules={[
        {
          required: true,
          message: "Please select your dog's image!",
        },
      ]}
    >
      <Upload
        name="dog-image"
        listType="picture-card"
        className="dog-image-uploader"
        maxCount={1}
        beforeUpload={beforeUpload}
        onRemove={() => setFileList([])}
      >
        {imageUrl ? (
          <img src={imageUrl} alt="new dog image" style={{ width: "100%" }} />
        ) : (
          fileList.length === 0 && imageUploadButton
        )}
      </Upload>
    </FormItem>
  );

  const nameFormItem = (
    <FormItem
      label="Name"
      name="name"
      rules={[
        {
          required: true,
          message: "Please input your dog's name!",
        },
      ]}
    >
      <Input placeholder="Goodest boi" />
    </FormItem>
  );

  const breedFormItem = (
    <FormItem
      label="Breed"
      name="breed"
      rules={[
        {
          required: true,
          min: 1,
          message: "Please input your dog's breed!",
        },
      ]}
    >
      <Input placeholder="Chihuahua" />
    </FormItem>
  );

  const birthYearFormItem = (
    <FormItem
      label="Birth Year"
      name="birth_year"
      rules={[
        {
          required: true,
          message: "Please input your dog's birth year!",
        },
      ]}
    >
      <DatePicker picker="year" />
    </FormItem>
  );

  const messageFormItem = (
    <FormItem label="Your message" name="message">
      <Input.TextArea />
    </FormItem>
  );

  return (
    <Modal
      open={visible}
      title="New Dog Post"
      okText="Submit"
      onCancel={onCancel}
      onOk={() => form.submit()}
      okButtonProps={{ loading }}
    >
      <Form
        form={form}
        layout="vertical"
        onFinishFailed={onFinishFailed}
        onFinish={onFinish}
      >
        {imageFormItem}
        {nameFormItem}
        {breedFormItem}
        {birthYearFormItem}
        {messageFormItem}
      </Form>
    </Modal>
  );
};

export default DogForm;
