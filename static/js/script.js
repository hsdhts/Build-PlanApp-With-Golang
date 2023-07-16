// Toggles line-through style for completed todos
document.addEventListener('DOMContentLoaded', function() {
    const todoList = document.querySelector('ul');
    todoList.addEventListener('click', function(event) {
        if (event.target.tagName === 'LI') {
            event.target.classList.toggle('completed');
        }
    });
});
