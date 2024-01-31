import { IsInt, IsNotEmpty, IsPositive, IsString, MaxLength } from "class-validator";

export class CreateAssetDto {
    @IsPositive()
    @IsInt()
    @IsNotEmpty()
    id: number;

    @MaxLength(255)
    @IsString()
    @IsNotEmpty()
    symbol: string;
}
