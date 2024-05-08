"use client";
import React from 'react';
import {Button, Input, Flex, Form} from 'antd';


const Home = () => {

    const [text, setText] = React.useState('');

    const handleChange = (event:  React.ChangeEvent<HTMLInputElement>) => {
        setText(event.target.value);
    }

    const handleSubmit = async () => {
        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}` + '/api/v1/string', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    value: text,
                }),
            }, )
            if(response.ok) {
                const data = await response.json();
                document.querySelector(".Label")?.insertAdjacentHTML('beforeend', newHtmlText(data.data.value));
            } else {
                console.log('Response not OK:', response);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    const newHtmlText = (text: string) => {
        return `<p>${text}</p>`;
    }

    return (
        <Flex gap="middle" vertical>
            <div className={"Label"}>
                Welcome to React
            </div>
            <Form className={"InputForm"} onSubmitCapture={handleSubmit}>
                <Input value={text} onChange={handleChange}/>
                <Button onClick={handleSubmit}>Submit</Button>
            </Form>
        </Flex>
    );
}

export default Home;