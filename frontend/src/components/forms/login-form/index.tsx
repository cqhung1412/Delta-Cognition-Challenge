import { Form, Input, Modal } from "antd";
import { useState } from "react";

const { Item: FormItem } = Form;

const LoginForm: React.FC<any> = ({
  onLogin,
  onFinishFailed,
  onCancel,
  visible,
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const onFinish = (values: any) => {
    setLoading(true);
    onLogin(values, () => {
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
      title="Login"
      okText="Login"
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

export default LoginForm;
