const API_BASE_URL = '/api';

// DOM Elements
const expenseForm = document.getElementById('expenseForm');
const expensesList = document.getElementById('expensesList');
const categoryFilter = document.getElementById('categoryFilter');
const sortOption = document.getElementById('sortOption');
const refreshBtn = document.getElementById('refreshBtn');
const totalAmount = document.getElementById('totalAmount');
const errorMessage = document.getElementById('errorMessage');
const successMessage = document.getElementById('successMessage');
const loadingMessage = document.getElementById('loadingMessage');
const submitBtn = document.getElementById('submitBtn');

// Set today's date as default
document.getElementById('date').valueAsDate = new Date();

// Load expenses on page load
document.addEventListener('DOMContentLoaded', () => {
    loadExpenses();
});

// Form submission
expenseForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    // Disable submit button to prevent multiple submissions
    submitBtn.disabled = true;
    submitBtn.textContent = 'Adding...';
    
    hideMessages();
    
    const formData = {
        amount: document.getElementById('amount').value.trim(),
        category: document.getElementById('category').value.trim(),
        description: document.getElementById('description').value.trim(),
        date: document.getElementById('date').value
    };
    
    try {
        const response = await fetch(`${API_BASE_URL}/expenses`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });
        
        const data = await response.json();
        
        if (response.ok) {
            showSuccess('Expense added successfully!');
            expenseForm.reset();
            document.getElementById('date').valueAsDate = new Date();
            loadExpenses();
        } else {
            showError(data.error || 'Failed to add expense');
        }
    } catch (error) {
        showError('Network error: ' + error.message);
    } finally {
        submitBtn.disabled = false;
        submitBtn.textContent = 'Add Expense';
    }
});

// Filter and sort changes
categoryFilter.addEventListener('change', loadExpenses);
sortOption.addEventListener('change', loadExpenses);
refreshBtn.addEventListener('click', loadExpenses);

// Load expenses from API
async function loadExpenses() {
    loadingMessage.style.display = 'block';
    expensesList.innerHTML = '';
    
    try {
        const category = categoryFilter.value;
        const sort = sortOption.value;
        
        let url = `${API_BASE_URL}/expenses?`;
        if (category) url += `category=${encodeURIComponent(category)}&`;
        if (sort) url += `sort=${encodeURIComponent(sort)}`;
        
        const response = await fetch(url);
        const expenses = await response.json();
        
        loadingMessage.style.display = 'none';
        
        if (expenses.length === 0) {
            expensesList.innerHTML = '<div class="empty-message">No expenses found</div>';
            totalAmount.textContent = '0.00';
            return;
        }
        
        displayExpenses(expenses);
        updateTotal(expenses);
        updateCategoryFilter(expenses);
    } catch (error) {
        loadingMessage.style.display = 'none';
        expensesList.innerHTML = `<div class="error-message">Error loading expenses: ${error.message}</div>`;
    }
}

// Display expenses in table
function displayExpenses(expenses) {
    const table = `
        <table>
            <thead>
                <tr>
                    <th>Date</th>
                    <th>Category</th>
                    <th>Description</th>
                    <th>Amount (₹)</th>
                </tr>
            </thead>
            <tbody>
                ${expenses.map(expense => `
                    <tr>
                        <td>${formatDate(expense.date)}</td>
                        <td>${escapeHtml(expense.category)}</td>
                        <td>${escapeHtml(expense.description)}</td>
                        <td>${parseFloat(expense.amount).toFixed(2)}</td>
                    </tr>
                `).join('')}
            </tbody>
        </table>
    `;
    
    expensesList.innerHTML = table;
}

// Calculate and update total
function updateTotal(expenses) {
    const total = expenses.reduce((sum, expense) => {
        return sum + parseFloat(expense.amount);
    }, 0);
    
    totalAmount.textContent = total.toFixed(2);
}

// Update category filter options
function updateCategoryFilter(expenses) {
    const categories = [...new Set(expenses.map(e => e.category))].sort();
    const currentValue = categoryFilter.value;
    
    categoryFilter.innerHTML = '<option value="">All Categories</option>';
    categories.forEach(category => {
        const option = document.createElement('option');
        option.value = category;
        option.textContent = category;
        if (category === currentValue) {
            option.selected = true;
        }
        categoryFilter.appendChild(option);
    });
}

// Format date for display
function formatDate(dateString) {
    const date = new Date(dateString + 'T00:00:00');
    return date.toLocaleDateString('en-IN', {
        year: 'numeric',
        month: 'short',
        day: 'numeric'
    });
}

// Escape HTML to prevent XSS
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Show error message
function showError(message) {
    errorMessage.textContent = message;
    errorMessage.style.display = 'block';
    successMessage.style.display = 'none';
    setTimeout(hideMessages, 5000);
}

// Show success message
function showSuccess(message) {
    successMessage.textContent = message;
    successMessage.style.display = 'block';
    errorMessage.style.display = 'none';
    setTimeout(hideMessages, 3000);
}

// Hide messages
function hideMessages() {
    errorMessage.style.display = 'none';
    successMessage.style.display = 'none';
}
