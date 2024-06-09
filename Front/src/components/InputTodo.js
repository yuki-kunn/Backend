import React, { useState, useEffect } from 'react';

export const InputTodo = (props) => {
  const { onClickAdd } = props;
  const [todoText, setTodoText] = useState('');

  // 初回レンダー時にローカルストレージから値を読み込む
  useEffect(() => {
    const savedTodoText = localStorage.getItem('todoText') || '';
    setTodoText(savedTodoText);
  }, []);

  // todoTextが変更されるたびにローカルストレージに保存する
  useEffect(() => {
    localStorage.setItem('todoText', todoText);
  }, [todoText]);

  const onChange = (event) => {
    setTodoText(event.target.value);
  };

  return (
    <div className="inputArea">
      <input
        placeholder="todoを入力"
        value={todoText}
        onChange={onChange}
      />
      <button onClick={() => {
        onClickAdd(todoText);
        setTodoText(''); // 追加後に入力フィールドをクリア
      }}>追加</button>
    </div>
  );
};