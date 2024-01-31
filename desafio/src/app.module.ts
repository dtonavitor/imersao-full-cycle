import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AssetsModule } from './assets/assets.module';
import { OrdersModule } from './orders/orders.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Asset } from './assets/entities/asset.entity';
import { Order } from './orders/entities/order.entity';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: "sqlite",
      database: "./src/database/database.sqlite",
      synchronize: true,
      logging: true,
      entities: [Order, Asset]
    }),
    AssetsModule, 
    OrdersModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
