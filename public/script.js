document.addEventListener('DOMContentLoaded', () => {
    const authContainer = document.getElementById('auth-container');
    const todoContainer = document.getElementById('todo-container');
    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');
    const todoForm = document.getElementById('todo-form');
    const todoList = document.getElementById('todo-list');
    const logoutBtn = document.getElementById('logout-btn');
    const todoItemTemplate = document.getElementById('todo-item-template');

    const apiUrl = 'http://127.0.0.1:8080/v1';
    let token = localStorage.getItem('token');

    const showAuth = () => {
        authContainer.classList.remove('hidden');
        todoContainer.classList.add('hidden');
    };

    const showTodos = () => {
        authContainer.classList.add('hidden');
        todoContainer.classList.remove('hidden');
        fetchTodos();
    };

    const fetchTodos = async () => {
        try {
            const response = await fetch(`${apiUrl}/todos/`,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            if (!response.ok) {
                throw new Error('Failed to fetch todos');
            }
            const result = await response.json();
            renderTodos(result.data);
        } catch (error) {
            console.error('Error fetching todos:', error);
            showAuth();
        }
    };

    const renderTodos = (todos) => {
        todoList.innerHTML = '';
        if (todos) {
             todos.forEach(todo => {
                // Clone the template for a new todo item
                const todoItem = todoItemTemplate.content.cloneNode(true);

                // Get the elements from the template
                const titleSpan = todoItem.querySelector('.todo-title');
                const statusSpan = todoItem.querySelector('.todo-status');
                const deleteBtn = todoItem.querySelector('.delete-btn');

                const editBtn = todoItem.querySelector('.edit-btn');
                const toggleStatusBtn = todoItem.querySelector('.toggle-status-btn');

                // Populate the data
                titleSpan.textContent = todo.title;
                statusSpan.textContent = `Status: ${todo.status}`;
                
                // Add event listener for the delete button
                deleteBtn.addEventListener('click', () => {
                    deleteTodo(todo.id);
                });

                // Add event listener for the edit button
                editBtn.addEventListener('click', () => {
                    window.location.href = `/public/edit.html?id=${todo.id}`;
                });

                // Add event listener for the toggle status button
                toggleStatusBtn.addEventListener('click', () => {
                    toggleTodoStatus(todo.id, todo.status);
                });

                // Append the new item to the list
                todoList.appendChild(todoItem);
            });
        }
    };

    const toggleTodoStatus = async (todoId, currentStatus) => {
        const newStatus = currentStatus === 'pending' ? 'completed' : 'pending';
        try {
            const response = await fetch(`${apiUrl}/todos/${todoId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ status: newStatus })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || 'Failed to update todo');
            }

            fetchTodos();
        } catch (error) {
            console.error('Error updating todo:', error);
            alert(error.message);
        }
    };

    const deleteTodo = async (todoId) => {
        try {
            const response = await fetch(`${apiUrl}/todos/${todoId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
    
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || 'Failed to delete todo');
            }
    
            // If deletion is successful, re-fetch and render the todos
            fetchTodos(); 
        } catch (error) {
            console.error('Error deleting todo:', error);
            alert(error.message);
        }
    }

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const username = document.getElementById('login-username').value;
        const password = document.getElementById('login-password').value;

        try {
            const response = await fetch(`${apiUrl}/auth/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const result = await response.json();
            token = result.data.token;
            localStorage.setItem('token', token);
            showTodos();
        } catch (error) {
            console.error('Login error:', error);
            alert('Login failed. Please check your credentials.');
        }
    });

    registerForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const username = document.getElementById('register-username').value;
        const password = document.getElementById('register-password').value;

        try {
            const response = await fetch(`${apiUrl}/auth/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                throw new Error('Registration failed');
            }

            alert('Registration successful! Please login.');
            registerForm.reset();
        } catch (error) {
            console.error('Registration error:', error);
            alert('Registration failed. Please try again.');
        }
    });

    todoForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const title = document.getElementById('todo-title').value;

        try {
            const response = await fetch(`${apiUrl}/todos/`,
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ title })
            });

            if (!response.ok) {
                throw new Error('Failed to add todo');
            }

            fetchTodos();
            todoForm.reset();
        } catch (error) {
            console.error('Error adding todo:', error);
        }
    });

    logoutBtn.addEventListener('click', () => {
        token = null;
        localStorage.removeItem('token');
        showAuth();
    });

    if (token) {
        showTodos();
    } else {
        showAuth();
    }
});
