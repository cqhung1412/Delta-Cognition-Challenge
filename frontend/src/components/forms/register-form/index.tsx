import { Form, Input, Modal } from "antd";
import { useState } from "react";

const { Item: FormItem } = Form;

const RegisterForm: React.FC<any> = ({
  onRegister,
  onFinishFailed,
  onCancel,
  visible,
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const onFinish = (values: any) => {
    setLoading(true);
    onRegister(values, () => {
      setLoading(false);
    });
  };

  const emailFormItem = (
    <FormItem
      label="Email"
      name="email"
      rules={[
        {
          required: true,
          type: "email",
          message: "Please input your email!",
        },
      ]}
    >
      <Input placeholder="example@gmail.com" />
    </FormItem>
  );

  const passwordFormItem = (
    <FormItem
      label="Password"
      name="password"
      rules={[
        {
          required: true,
          min: 6,
          message: "Please input your password!",
        },
      ]}
    >
      <Input.Password placeholder="*******" />
    </FormItem>
  );

  return (
    <Modal
      open={visible}
      title="Register"
      okText="Register"
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
        {emailFormItem}
        {passwordFormItem}
      </Form>
    </Modal>
  );
};

export default RegisterForm;
