import React, { useEffect, useState } from "react";
import { Row, Col, Typography, Tooltip, Space, Image, Button } from "antd";
import { HeartTwoTone, LinkOutlined } from "@ant-design/icons";
import axios from "../../axios";
import "./index.css";

const Feed: React.FC<{ isAuth: boolean }> = ({ isAuth }) => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    listDogPosts();
  }, []);

  const listDogPosts = () => {
    axios
      .get("/dogs?limit=20&offset=0")
      .then((res) => {
        console.log(res);
        setPosts(res.data || []);
      })
      .catch((err) => console.log(err));
  };

  const listSimilarDogs = (dogId: number) => {
    axios
      .get(`/recommend/dogs?dog_id=${dogId}&limit=20&offset=0`)
      .then((res) => {
        console.log(res);
        setPosts(res.data || []);
      })
      .catch((err) => console.log(err));
  };

  return (
    <div className="feed-container">
      <Space
        direction="vertical"
        size="middle"
        style={{ display: "flex" }}
      ></Space>
      {posts.map((post: any) => {
        const { dog, photo_url } = post;
        return (
          <Row gutter={12} key={dog.id} className="post-container">
            <Col span={24}>
              <div className="card">
                <Image alt={dog.name} src={photo_url} />
                <div className="container">
                  <Typography.Title level={4}>
                    <b>
                      {dog.name} ({dog.breed})
                    </b>
                  </Typography.Title>
                  <Typography.Text>
                    {dog.message.String || "No message"}
                  </Typography.Text>
                </div>
                <div className="buttons">
                  <Tooltip title="Like">
                    <Button
                      type="text"
                      icon={<HeartTwoTone twoToneColor="#eb2f96" />}
                    />
                  </Tooltip>
                  <Tooltip title="View similar">
                    <Button type="text" icon={<LinkOutlined />} onClick={() => listSimilarDogs(dog.id)} />
                  </Tooltip>
                </div>
              </div>
            </Col>
          </Row>
        );
      })}
    </div>
  );
};

export default Feed;
