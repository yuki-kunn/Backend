import React, { useState, useEffect } from 'react';
import { InputTodo } from './components/InputTodo';
import { IncompleteTodos } from './components/IncompleteTodos';
import { CompleteTodos } from './components/CompleteTodos';

const TodoApp = () => {
  const [incompleteTodos, setIncompleteTodos] = useState([]);
  const [completeTodos, setCompleteTodos] = useState([]);

  // 初回レンダー時にローカルストレージからデータを読み込む
  useEffect(() => {
    const savedIncompleteTodos = JSON.parse(localStorage.getItem('incompleteTodos')) || [];
    const savedCompleteTodos = JSON.parse(localStorage.getItem('completeTodos')) || [];
    setIncompleteTodos(savedIncompleteTodos);
    setCompleteTodos(savedCompleteTodos);
  }, []);

  // incompleteTodosまたはcompleteTodosが変更されるたびにローカルストレージを更新する
  useEffect(() => {
    localStorage.setItem('incompleteTodos', JSON.stringify(incompleteTodos));
    localStorage.setItem('completeTodos', JSON.stringify(completeTodos));
  }, [incompleteTodos, completeTodos]);

  const onClickAdd = (todoText) => {
    if (todoText === '') return;
    const newTodos = [...incompleteTodos, todoText];
    setIncompleteTodos(newTodos);
  };

  const onClickComplete = (index) => {
    const newIncompleteTodos = [...incompleteTodos];
    const completeTodo = newIncompleteTodos.splice(index, 1)[0];
    setIncompleteTodos(newIncompleteTodos);
    setCompleteTodos([...completeTodos, completeTodo]);
  };

  const onClickDelete = (index) => {
    const newIncompleteTodos = [...incompleteTodos];
    newIncompleteTodos.splice(index, 1);
    setIncompleteTodos(newIncompleteTodos);
  };

  const onClickBack = (index) => {
    const newCompleteTodos = [...completeTodos];
    const backTodo = newCompleteTodos.splice(index, 1)[0];
    setCompleteTodos(newCompleteTodos);
    setIncompleteTodos([...incompleteTodos, backTodo]);
  };

  return (
    <div className="outline">
      <InputTodo onClickAdd={onClickAdd} />
      <IncompleteTodos
        todos={incompleteTodos}
        onClickComplete={onClickComplete}
        onClickDelete={onClickDelete}
      />
      <CompleteTodos
        todos={completeTodos}
        onClickBack={onClickBack}
      />
    </div>
  );
};

export default TodoApp;