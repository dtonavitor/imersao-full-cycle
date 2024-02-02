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
      let asset = await this.assetRepo.findOneBy({
        id: createOrderDto.asset_id,
      });

      
      if (!asset && !createOrderDto.symbol) {
        return {
          message: 'Asset não encontrado. Informe o símbolo do ativo.',
          status: 404,
        };
      } else if (!asset && createOrderDto.symbol) {
        asset = this.assetRepo.create({
          id: createOrderDto.asset_id,
          symbol: createOrderDto.symbol,
        });
      }

      const order = this.orderRepo.create({
        ...createOrderDto,
        asset: asset,
      });

      await this.orderRepo.save(order);
      return order;
    } catch (error) {
      console.log(error);
    }
  }

  findAll() {
    return this.orderRepo.find();
  }

  findOne(id: string) {
    return this.orderRepo.findOneByOrFail({ id });
  }
}
