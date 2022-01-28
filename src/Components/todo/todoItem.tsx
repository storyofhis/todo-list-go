import { Todo } from "../../entities/Todo";

export const TodoItem: React.FC<Todo> = ({ title, description, isCompleted }) => {
  return (
    <article>
      <h3>{title}</h3>
      <p>{description}</p>
      <input type="checkbox" defaultChecked={isCompleted} />
    </article>
  );
};
