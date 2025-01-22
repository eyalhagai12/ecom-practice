interface Product {
    id: string;
    name: string;
    description: string;
    price: number;
    stock: number;
    image: string;
}

export interface StoreStats {
    plCash: number;
    plPercentage: number;
    totalSales: number;
    totalOrders: number;
    totalProducts: number;
}

export interface Store {
    id: string;
    name: string;
    description: string;
    products: Product[];
    stats: StoreStats;
}
