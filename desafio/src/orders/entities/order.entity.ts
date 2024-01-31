import { Asset } from "src/assets/entities/asset.entity";
import { Column, Entity, JoinColumn, OneToOne, PrimaryGeneratedColumn } from "typeorm";

export enum OrderStatus {
    OPEN = 'open',
    PENDING = 'pending',
    CLOSED = 'closed',
}

@Entity()
export class Order {
    @PrimaryGeneratedColumn("uuid")
    id: string;

    @OneToOne(() => Asset, { eager: true, cascade: ['insert'] })
    @JoinColumn({ name: 'asset_id' })
    asset: Asset;

    @Column({ type: 'decimal', precision: 10, scale: 2 })
    price: number;

    @Column()
    status: OrderStatus = OrderStatus.OPEN;
}
