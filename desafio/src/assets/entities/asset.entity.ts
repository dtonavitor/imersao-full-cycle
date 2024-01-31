import { Column, Entity, PrimaryColumn } from "typeorm";

@Entity()
export class Asset {
    @PrimaryColumn()
    id: number;

    @Column()
    symbol: string;
}
