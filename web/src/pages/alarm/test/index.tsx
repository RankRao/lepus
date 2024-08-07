import { PlusOutlined, FormOutlined, DeleteOutlined } from '@ant-design/icons';
import { Button, message, Input, Form, Card, Row, Col, Alert } from 'antd';
import React, { useState, useRef } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import { useAccess } from 'umi';

/* eslint-disable no-template-curly-in-string */
const validateMessages = {
  required: '${label}是必填项!',
  types: {
    email: '${label} is not a valid email!',
    number: '${label} is not a valid number!',
  },
  number: {
    range: '${label} must be between ${min} and ${max}',
  },
};

const sendTest: React.FC<{}> = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const actionRef = useRef<ActionType>();
  const access = useAccess();

  const [form] = Form.useForm();

  const [formValues, setFormValues] = useState({
    email_list: "",
    sms_list: "",
    phone_list: "",
    wechat_list: "",
    weburl: "",
  });


  //表单提交查询执行请求
  const asyncFetch = (values: {}, sendUrl: string) => {
    console.info(values);
    setLoading(true);
    const params = { ...values, };
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    fetch(sendUrl, {
      method: 'post',
      headers: headers,
      body: JSON.stringify(params),
    })
      .then((response) => response.json())
      .then((json) => {
        console.info(json.msg);
        setLoading(false);
        if (json.success == true) {
          message.success(json.msg);
        } else {
          message.error(json.msg);
        }

      })
      .catch((error) => {
        setLoading(false);
        console.log('post data failed', error);
      });
  };


  const onFinish = (fieldValue: []) => {
    const values = {
      email_list: fieldValue["email_list"],
    };
    setFormValues(values);
    asyncFetch(values, '/api/v1/alarm/test/send_email');
  };

  return (
    <PageContainer >
      <Row gutter={[16, 24]} style={{ marginTop: '0px' }}>

        <Col span={8}>
          <Card title="发送邮件测试" bordered={false} >
            <Form name="nest-messages" layout="vertical" onFinish={onFinish} >
              <Form.Item name="email_list" label="收件人邮箱" initialValue={formValues.email_list} tooltip={"发送前确保邮件网关配置正确，多个收件人使用英文分号分隔"} rules={[{ required: true }]}>
                <Input />
              </Form.Item>
              <Form.Item wrapperCol={{ offset: 8 }}>
                <Button type="primary" htmlType="submit" loading={loading}>
                  发送邮件
                </Button>
              </Form.Item>
            </Form>
          </Card>
        </Col>


      </Row>
    </PageContainer>
  );
};

export default sendTest;
