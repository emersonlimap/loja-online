-- Script de inicialização do banco de dados
-- Este script é executado automaticamente quando o container PostgreSQL é criado

-- Criar extensões úteis
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Inserir usuário admin padrão
-- Senha: leozin@123 (hash bcrypt)
INSERT INTO users (name, email, password, role, active, created_at, updated_at, products, customers, inventory, sales, reports)
SELECT 
    'Administrador',
    'leozinsurfwear@gmail.com',
    '$2a$10$YourHashWillBeReplacedAutomatically', -- Hash será substituído no main.go
    'admin',
    true,
    NOW(),
    NOW(),
    true,  -- Permissão para produtos
    true,  -- Permissão para clientes
    true,  -- Permissão para inventário
    true,  -- Permissão para vendas
    true   -- Permissão para relatórios
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE email = 'leozinsurfwear@gmail.com'
);

-- Inserir alguns produtos de exemplo
INSERT INTO products (name, description, category, brand, price, cost_price, sku, color, size, material, gender, season, active, image_url, created_at, updated_at)
SELECT * FROM (VALUES
    ('Camiseta Básica Branca', 'Camiseta básica em algodão 100%', 'Camiseta', 'LeoZin Surfwear', 49.90, 25.00, 'CAM-BAS-BCO-P', 'Branco', 'P', 'Algodão', 'Unissex', 'Verão', true, '', NOW(), NOW()),
    ('Camiseta Básica Branca', 'Camiseta básica em algodão 100%', 'Camiseta', 'LeoZin Surfwear', 49.90, 25.00, 'CAM-BAS-BCO-M', 'Branco', 'M', 'Algodão', 'Unissex', 'Verão', true, '', NOW(), NOW()),
    ('Camiseta Básica Branca', 'Camiseta básica em algodão 100%', 'Camiseta', 'LeoZin Surfwear', 49.90, 25.00, 'CAM-BAS-BCO-G', 'Branco', 'G', 'Algodão', 'Unissex', 'Verão', true, '', NOW(), NOW()),
    ('Bermuda Surf Azul', 'Bermuda para surf em tecido quick-dry', 'Bermuda', 'LeoZin Surfwear', 89.90, 45.00, 'BER-SUR-AZU-P', 'Azul', 'P', 'Poliéster', 'Masculino', 'Verão', true, '', NOW(), NOW()),
    ('Bermuda Surf Azul', 'Bermuda para surf em tecido quick-dry', 'Bermuda', 'LeoZin Surfwear', 89.90, 45.00, 'BER-SUR-AZU-M', 'Azul', 'M', 'Poliéster', 'Masculino', 'Verão', true, '', NOW(), NOW()),
    ('Top Surf Feminino', 'Top para surf com proteção UV', 'Top', 'LeoZin Surfwear', 79.90, 40.00, 'TOP-SUR-PTO-P', 'Preto', 'P', 'Lycra', 'Feminino', 'Verão', true, '', NOW(), NOW()),
    ('Top Surf Feminino', 'Top para surf com proteção UV', 'Top', 'LeoZin Surfwear', 79.90, 40.00, 'TOP-SUR-PTO-M', 'Preto', 'M', 'Lycra', 'Feminino', 'Verão', true, '', NOW(), NOW())
) AS v(name, description, category, brand, price, cost_price, sku, color, size, material, gender, season, active, image_url, created_at, updated_at)
WHERE NOT EXISTS (SELECT 1 FROM products WHERE sku = v.sku);

-- Inserir estoque inicial para os produtos
INSERT INTO inventory_items (product_id, quantity, min_stock, max_stock, location, created_at, updated_at)
SELECT p.id, 50, 10, 200, 'Estoque Principal', NOW(), NOW()
FROM products p
WHERE NOT EXISTS (
    SELECT 1 FROM inventory_items i WHERE i.product_id = p.id
);

-- Inserir cliente de exemplo
INSERT INTO customers (name, email, phone, cpf, gender, active, created_at, updated_at)
SELECT 
    'Cliente Exemplo',
    'cliente@exemplo.com',
    '(11) 99999-9999',
    '123.456.789-00',
    'Masculino',
    true,
    NOW(),
    NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM customers WHERE email = 'cliente@exemplo.com'
);

-- Mensagem de confirmação (será exibida nos logs do container)
DO $$
BEGIN
    RAISE NOTICE 'Base de dados inicializada com sucesso!';
    RAISE NOTICE 'Usuário admin criado: leozinsurfwear@gmail.com / leozin@123';
    RAISE NOTICE 'Produtos de exemplo inseridos';
    RAISE NOTICE 'Estoque inicial configurado';
END $$;