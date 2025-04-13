// カード種類の定義
export type CardType = 'rakuten' | 'mufg' | 'epos';

// 表示用のカード名マッピング
export const cardTypeDisplayNames: Record<CardType, string> = {
  'rakuten': '楽天カード',
  'mufg': 'MUFG DCカード',
  'epos': 'EPOSカード'
};
