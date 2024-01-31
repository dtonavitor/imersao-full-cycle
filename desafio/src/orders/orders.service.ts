import { Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Asset } from '../assets/entities/asset.entity';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectRepository(Order) private orderRepo: Repository<Order>,
    @InjectRepository(Asset) private assetRepo: Repository<Asset>,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    try {
      const asset = await this.assetRepo.findOneByOrFail({
        id: createOrderDto.asset_id,
      });

      const order = this.orderRepo.create({
        ...createOrderDto,
        asset: asset,
      });

      return this.orderRepo.save(order);
    } catch (error) {
        return {
          "message": 'Asset não encontrado',
          "status": 404,
        }
    }
  }

  findAll() {
    return this.orderRepo.find();
  }

  findOne(id: string) {
    return this.orderRepo.findOneByOrFail({ id });
  }
}
