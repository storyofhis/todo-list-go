import { Todo } from "../../entities/Todo";
import { TodoItem } from "../todo/todoItem";
import { Heading, Avatar, Box, Center, Text, Stack, Button, Link, Badge, useColorModeValue } from "@chakra-ui/react";

type Props = {
  todos: Todo[];
};

export const TodoList: React.FC<Props> = ({ todos }) => {
  return (
    <Center py={6}>
      <ul>
        {todos.map((todo, index) => (
          <Box maxW={"320px"} w={"full"} boxShadow={"2xl"} rounded={"lg"} p={6} textAlign={"center"} key={index}>
            <TodoItem title={todo.title} description={todo.description} isCompleted={todo.isCompleted} />
          </Box>
        ))}
      </ul>
    </Center>
  );
};
