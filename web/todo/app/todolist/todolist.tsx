
export interface TodoItem {
    id: number;
    descr: string;
}
export interface TodoList {
    id: number;
    name: string;
    items: TodoItem[];
}

export function TodoListBox(props: {}) {

}