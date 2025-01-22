import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Store } from "@/interfaces/store";

export default function Home() {
  const stores: Store[] = [
    {
      id: "1",
      name: "Store 1",
      description: "Store 1 description",
      products: [
        {
          id: "1",
          name: "Product 1",
          description: "Product 1 description",
          price: 10,
          stock: 100,
          image: "https://via.placeholder.com/150",
        },
      ],
      stats: {
        plCash: 1000,
        plPercentage: 10,
        totalSales: 100,
        totalOrders: 10,
        totalProducts: 1,
      },
    },
    {
      id: "2",
      name: "Store 2",
      description: "Store 2 description",
      products: [
        {
          id: "1",
          name: "Product 1",
          description: "Product 1 description",
          price: 10,
          stock: 100,
          image: "https://via.placeholder.com/150",
        },
        {
          id: "2",
          name: "Product 2",
          description: "Product 2 description",
          price: 20,
          stock: 50,
          image: "https://via.placeholder.com/150",
        },
      ],
      stats: {
        plCash: -2000,
        plPercentage: -20,
        totalSales: 200,
        totalOrders: 20,
        totalProducts: 2,
      },
    },
  ];

  return (
    <div>
      <h1 className="text-4xl font-bold">Stores</h1>
      <div>
        <Table>
          <TableHeader>
            <TableHead>Name</TableHead>
            <TableHead>Description</TableHead>
            <TableHead>Profit/Loss Cash</TableHead>
            <TableHead>Profit/Loss Percent</TableHead>
          </TableHeader>
          <TableBody>
          {stores.map((store) => (
            <TableRow key={store.id}>
              <TableCell>{store.name}</TableCell>
              <TableCell>{store.description}</TableCell>
              <TableCell className={store.stats.plCash > 0 ? 'text-green-700' : 'text-red-800'}>{store.stats.plCash}$</TableCell>
              <TableCell className={store.stats.plPercentage > 0 ? 'text-green-700' : 'text-red-800'}>{store.stats.plPercentage}%</TableCell>
            </TableRow>
          ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
