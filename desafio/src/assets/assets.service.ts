import { Injectable } from '@nestjs/common';
import { CreateAssetDto } from './dto/create-asset.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Asset } from './entities/asset.entity';
import { Repository } from 'typeorm';

@Injectable()
export class AssetsService {
  constructor(@InjectRepository(Asset) private assetRepo: Repository<Asset>) {}

  create(createAssetDto: CreateAssetDto) {
    const asset = this.assetRepo.create(createAssetDto);
    return this.assetRepo.save(asset);
  }

  findAll() {
    return this.assetRepo.find();
  }

  findOne(id: number) {
    return this.assetRepo.findOne({
      where: { id },
    });
  }
}
