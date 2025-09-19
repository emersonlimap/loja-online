class LojaOnlineApp {
    constructor() {
        this.baseURL = '/api/v1';
        this.token = localStorage.getItem('token');
        this.currentUser = null;
        this.init();
    }

    init() {
        this.setupEventListeners();
        this.checkAuth();
    }

    setupEventListeners() {
        // Login form
        const loginForm = document.getElementById('loginForm');
        if (loginForm) {
            loginForm.addEventListener('submit', this.handleLogin.bind(this));
        }

        // Navigation
        document.addEventListener('click', (e) => {
            if (e.target.matches('[data-action]')) {
                const action = e.target.getAttribute('data-action');
                this.handleAction(action, e.target);
            }
        });
    }

    async makeRequest(endpoint, options = {}) {
        const config = {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        };

        if (this.token) {
            config.headers.Authorization = `Bearer ${this.token}`;
        }

        try {
            const response = await fetch(`${this.baseURL}${endpoint}`, config);
            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || 'Erro na requisição');
            }

            return data;
        } catch (error) {
            console.error('Request error:', error);
            throw error;
        }
    }

    async handleLogin(e) {
        e.preventDefault();
        const formData = new FormData(e.target);
        const credentials = {
            email: formData.get('email'),
            password: formData.get('password')
        };

        try {
            const response = await this.makeRequest('/auth/login', {
                method: 'POST',
                body: JSON.stringify(credentials)
            });

            this.token = response.token;
            localStorage.setItem('token', this.token);
            this.currentUser = response.user;
            
            window.location.href = '/dashboard';
        } catch (error) {
            alert('Erro no login: ' + error.message);
        }
    }

    async checkAuth() {
        if (!this.token) {
            if (window.location.pathname !== '/login') {
                window.location.href = '/login';
            }
            return;
        }

        try {
            const user = await this.makeRequest('/auth/me');
            this.currentUser = user;
            this.updateUI();
        } catch (error) {
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
    }

    updateUI() {
        const userNameElements = document.querySelectorAll('[data-user-name]');
        userNameElements.forEach(el => {
            el.textContent = this.currentUser?.name || '';
        });

        // Show/hide menu items based on permissions
        if (this.currentUser) {
            this.updateMenuPermissions();
        }
    }

    updateMenuPermissions() {
        const permissions = this.currentUser.permissions;
        
        document.querySelectorAll('[data-permission]').forEach(element => {
            const requiredPermission = element.getAttribute('data-permission');
            if (!permissions[requiredPermission]) {
                element.style.display = 'none';
            }
        });
    }

    async handleAction(action, element) {
        switch (action) {
            case 'logout':
                this.logout();
                break;
            case 'load-products':
                await this.loadProducts();
                break;
            case 'load-customers':
                await this.loadCustomers();
                break;
            case 'load-sales':
                await this.loadSales();
                break;
            case 'load-inventory':
                await this.loadInventory();
                break;
            case 'load-reports':
                await this.loadReports();
                break;
        }
    }

    logout() {
        localStorage.removeItem('token');
        this.token = null;
        this.currentUser = null;
        window.location.href = '/login';
    }

    async loadProducts() {
        try {
            const products = await this.makeRequest('/products');
            this.renderProducts(products);
        } catch (error) {
            console.error('Erro ao carregar produtos:', error);
        }
    }

    async loadCustomers() {
        try {
            const customers = await this.makeRequest('/customers');
            this.renderCustomers(customers);
        } catch (error) {
            console.error('Erro ao carregar clientes:', error);
        }
    }

    async loadSales() {
        try {
            const sales = await this.makeRequest('/sales');
            this.renderSales(sales);
        } catch (error) {
            console.error('Erro ao carregar vendas:', error);
        }
    }

    async loadInventory() {
        try {
            const inventory = await this.makeRequest('/inventory');
            this.renderInventory(inventory);
        } catch (error) {
            console.error('Erro ao carregar estoque:', error);
        }
    }

    async loadReports() {
        try {
            const reports = await this.makeRequest('/reports/sales');
            this.renderReports(reports);
        } catch (error) {
            console.error('Erro ao carregar relatórios:', error);
        }
    }

    renderProducts(products) {
        const container = document.getElementById('productsContainer');
        if (!container) return;

        container.innerHTML = `
            <div class="card">
                <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 1rem;">
                    <h2>Produtos</h2>
                    <button class="btn btn-primary" onclick="showProductModal()">Novo Produto</button>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>SKU</th>
                            <th>Nome</th>
                            <th>Categoria</th>
                            <th>Preço</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${products.map(product => `
                            <tr>
                                <td>${product.sku}</td>
                                <td>${product.name}</td>
                                <td>${product.category}</td>
                                <td>R$ ${product.price.toFixed(2)}</td>
                                <td>
                                    <button class="btn btn-secondary" onclick="editProduct(${product.id})">Editar</button>
                                    <button class="btn btn-danger" onclick="deleteProduct(${product.id})">Excluir</button>
                                </td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;
    }

    renderCustomers(customers) {
        const container = document.getElementById('customersContainer');
        if (!container) return;

        container.innerHTML = `
            <div class="card">
                <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 1rem;">
                    <h2>Clientes</h2>
                    <button class="btn btn-primary" onclick="showCustomerModal()">Novo Cliente</button>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Email</th>
                            <th>CPF</th>
                            <th>Telefone</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${customers.map(customer => `
                            <tr>
                                <td>${customer.name}</td>
                                <td>${customer.email || '-'}</td>
                                <td>${customer.cpf}</td>
                                <td>${customer.phone || '-'}</td>
                                <td>
                                    <button class="btn btn-secondary" onclick="editCustomer(${customer.id})">Editar</button>
                                    <button class="btn btn-danger" onclick="deleteCustomer(${customer.id})">Excluir</button>
                                </td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;
    }

    renderSales(sales) {
        const container = document.getElementById('salesContainer');
        if (!container) return;

        container.innerHTML = `
            <div class="card">
                <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 1rem;">
                    <h2>Vendas</h2>
                    <button class="btn btn-primary" onclick="showSaleModal()">Nova Venda</button>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Cliente</th>
                            <th>Data</th>
                            <th>Total</th>
                            <th>Status</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${sales.map(sale => `
                            <tr>
                                <td>#${sale.id}</td>
                                <td>${sale.customer?.name || 'Cliente Avulso'}</td>
                                <td>${new Date(sale.sale_date).toLocaleDateString('pt-BR')}</td>
                                <td>R$ ${sale.final_amount.toFixed(2)}</td>
                                <td><span class="status-${sale.status}">${sale.status}</span></td>
                                <td>
                                    <button class="btn btn-secondary" onclick="viewSale(${sale.id})">Ver</button>
                                    <button class="btn btn-danger" onclick="cancelSale(${sale.id})">Cancelar</button>
                                </td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;
    }

    renderInventory(inventory) {
        const container = document.getElementById('inventoryContainer');
        if (!container) return;

        container.innerHTML = `
            <div class="card">
                <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 1rem;">
                    <h2>Controle de Estoque</h2>
                    <button class="btn btn-primary" onclick="showInventoryModal()">Ajustar Estoque</button>
                </div>
                <table>
                    <thead>
                        <tr>
                            <th>Produto</th>
                            <th>SKU</th>
                            <th>Quantidade</th>
                            <th>Mín.</th>
                            <th>Máx.</th>
                            <th>Status</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${inventory.map(item => {
                            let status = 'ok';
                            if (item.quantity <= item.min_stock) status = 'low';
                            if (item.quantity === 0) status = 'out';
                            
                            return `
                                <tr>
                                    <td>${item.product.name}</td>
                                    <td>${item.product.sku}</td>
                                    <td>${item.quantity}</td>
                                    <td>${item.min_stock}</td>
                                    <td>${item.max_stock}</td>
                                    <td><span class="stock-${status}">${status === 'ok' ? 'OK' : status === 'low' ? 'Baixo' : 'Zerado'}</span></td>
                                    <td>
                                        <button class="btn btn-secondary" onclick="adjustStock(${item.product_id})">Ajustar</button>
                                        <button class="btn btn-primary" onclick="viewMovements(${item.product_id})">Movimentos</button>
                                    </td>
                                </tr>
                            `;
                        }).join('')}
                    </tbody>
                </table>
            </div>
        `;
    }

    renderReports(reports) {
        const container = document.getElementById('reportsContainer');
        if (!container) return;

        container.innerHTML = `
            <div class="card">
                <h2>Relatórios de Vendas</h2>
                <div class="dashboard-stats">
                    <div class="stat-card">
                        <div class="stat-number">${reports.total_sales}</div>
                        <div class="stat-label">Total de Vendas</div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-number">R$ ${reports.total_amount.toFixed(2)}</div>
                        <div class="stat-label">Faturamento Total</div>
                    </div>
                </div>
                
                <h3>Produtos Mais Vendidos</h3>
                <table>
                    <thead>
                        <tr>
                            <th>Produto</th>
                            <th>Quantidade</th>
                            <th>Receita</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${reports.top_products.map(product => `
                            <tr>
                                <td>${product.product_name}</td>
                                <td>${product.quantity}</td>
                                <td>R$ ${product.revenue.toFixed(2)}</td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;
    }
}

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    window.app = new LojaOnlineApp();
});
