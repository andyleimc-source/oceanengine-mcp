# 巨量引擎 Marketing API 完整目录

SDK: `github.com/oceanengine/ad_open_sdk_go@v1.1.83`
总计: 1128 个 API
状态: [ ] 未实现  [x] 已实现

---

## 01. OAuth2 认证

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 01-001 | Oauth2AccessTokenApi | 获取 Access Token | [x] |
| 01-002 | Oauth2RefreshTokenApi | 刷新 Token | [x] |
| 01-003 | Oauth2RenewTokenApi | 续期 Token | [ ] |
| 01-004 | Oauth2AppAccessTokenApi | 获取应用级 Token | [ ] |
| 01-005 | Oauth2AdvertiserGetApi | 获取已授权账户列表 | [x] |

## 02. 广告主/账户管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 02-001 | AdvertiserInfoV2Api | 广告主信息查询 | [ ] |
| 02-002 | AdvertiserPublicInfoV2Api | 广告主公开信息 | [ ] |
| 02-003 | AdvertiserBudgetGetV2Api | 广告主预算查询 | [ ] |
| 02-004 | AdvertiserUpdateBudgetV2Api | 更新广告主预算 | [ ] |
| 02-005 | AdvertiserVerifyInfoGetV30Api | 广告主认证信息 | [ ] |
| 02-006 | AdvertiserAvatarGetV2Api | 获取广告主头像 | [ ] |
| 02-007 | AdvertiserAvatarSubmitV2Api | 提交广告主头像 | [ ] |
| 02-008 | AdvertiserAvatarUploadV2Api | 上传广告主头像 | [ ] |
| 02-009 | AdvertiserAttachmentUploadV30Api | 上传附件 | [ ] |
| 02-010 | AdvertiserDeliveryPkgConfigV30Api | 投放包配置 | [ ] |
| 02-011 | AdvertiserDeliveryPkgDeleteV30Api | 删除投放包 | [ ] |
| 02-012 | AdvertiserDeliveryPkgGetV30Api | 查询投放包 | [ ] |
| 02-013 | AdvertiserDeliveryPkgSubmitV30Api | 提交投放包 | [ ] |
| 02-014 | AdvertiserDeliveryQualificationDeleteV30Api | 删除投放资质 | [ ] |
| 02-015 | AdvertiserDeliveryQualificationListV30Api | 投放资质列表 | [ ] |
| 02-016 | AdvertiserDeliveryQualificationSubmitV30Api | 提交投放资质 | [ ] |
| 02-017 | AdvertiserQualificationCreateV2V2Api | 创建资质 | [ ] |
| 02-018 | AdvertiserQualificationGetV30Api | 查询资质 | [ ] |
| 02-019 | AdvertiserQualificationSelectV2V2Api | 筛选资质 | [ ] |
| 02-020 | AdvertiserQualificationSubmitV30Api | 提交资质 | [ ] |
| 02-021 | AccountFundGetV30Api | 账户资金查询 | [ ] |
| 02-022 | AccountUpdateV30Api | 账户更新 | [ ] |

## 03. 资金与财务

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 03-001 | AdvertiserFundGetV2Api | 广告主余额查询 | [ ] |
| 03-002 | AdvertiserFundDailyStatV2Api | 日流水统计 | [ ] |
| 03-003 | AdvertiserFundTransactionGetV2Api | 交易明细 | [ ] |
| 03-004 | AdvertiserFundGrantTransactionGetV2Api | 赠款交易明细 | [ ] |
| 03-005 | AdvertiserFundDetailGrantV2Api | 赠款明细 | [ ] |
| 03-006 | AdvertiserTransferableFundGetV2Api | 可转余额查询 | [ ] |
| 03-007 | FundSharedWalletBalanceGetV2Api | 共享钱包余额 | [ ] |
| 03-008 | ChargeListV30Api | 充值列表 | [ ] |
| 03-009 | ChargeResultV30Api | 充值结果 | [ ] |
| 03-010 | ChargeVerifyGetV30Api | 充值验证 | [ ] |
| 03-011 | DcdChargeSubmitV30Api | DCD充值提交 | [ ] |
| 03-012 | SvipChargeVerifyGetV30Api | SVIP充值验证 | [ ] |
| 03-013 | PrepayChargeGenerateFixRemiattanceCodeCreateV30Api | 生成固定汇款码 | [ ] |
| 03-014 | PrepayChargeGenerateRemittanceCodeCreateV30Api | 生成汇款码 | [ ] |
| 03-015 | RemittanceCodeListGetV30Api | 汇款码列表 | [ ] |
| 03-016 | RemittanceCodeListV30Api | 汇款码列表v2 | [ ] |
| 03-017 | FixRemittanceCodeListGetV30Api | 固定汇款码列表 | [ ] |
| 03-018 | WalletChargeVerifyGetV30Api | 钱包充值验证 | [ ] |
| 03-019 | WalletPrepayChargeGenerateRemittanceCodeCreateV30Api | 钱包汇款码 | [ ] |
| 03-020 | WalletRemittanceCodeListGetV30Api | 钱包汇款码列表 | [ ] |

## 04. 转账

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 04-001 | CgTransferCanTransferBalanceGetV30Api | 可转余额查询 | [ ] |
| 04-002 | CgTransferCanTransferTargetListV30Api | 可转目标列表 | [ ] |
| 04-003 | CgTransferCreateTransferV30Api | 创建转账 | [ ] |
| 04-004 | CgTransferQueryCanTransferBalanceV30Api | 查询可转余额 | [ ] |
| 04-005 | CgTransferQueryTransferBalanceV30Api | 查询转账余额 | [ ] |
| 04-006 | CgTransferQueryTransferDetailV30Api | 转账明细 | [ ] |
| 04-007 | CgTransferTransferBalanceGetV30Api | 转账余额 | [ ] |
| 04-008 | CgTransferTransferCreateV30Api | 创建转账v2 | [ ] |
| 04-009 | CgTransferTransferDetailGetV30Api | 转账详情 | [ ] |
| 04-010 | CgTransferWalletTransferCanTransferBalanceV30Api | 钱包可转余额 | [ ] |
| 04-011 | CgTransferWalletTransferCreateV30Api | 钱包转账 | [ ] |
| 04-012 | CgTransferWalletTransferDetailV30Api | 钱包转账明细 | [ ] |
| 04-013 | CgTransferWalletTransferListV30Api | 钱包转账列表 | [ ] |
| 04-014 | CustomerCenterFundTransferSeqCreateV2Api | 客户中心转账创建 | [ ] |
| 04-015 | CustomerCenterFundTransferSeqCommitV2Api | 客户中心转账提交 | [ ] |

## 05. 共享钱包

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 05-001 | SharedWalletAccountRelationGetV30Api | 账户关联查询 | [ ] |
| 05-002 | SharedWalletBudgetGetV30Api | 预算查询 | [ ] |
| 05-003 | SharedWalletBudgetSubmitV30Api | 预算提交 | [ ] |
| 05-004 | SharedWalletDailyStatGetV30Api | 日流水 | [ ] |
| 05-005 | SharedWalletMainWalletGetV30Api | 主钱包查询 | [ ] |
| 05-006 | SharedWalletSharedRelationCreateV30Api | 创建共享关系 | [ ] |
| 05-007 | SharedWalletSubWalletCreateV30Api | 创建子钱包 | [ ] |
| 05-008 | SharedWalletTransactionDetailGetV30Api | 交易明细 | [ ] |
| 05-009 | SharedWalletWalletAdvOperationLogGetV30Api | 广告操作日志 | [ ] |
| 05-010 | SharedWalletWalletBalanceGetV30Api | 钱包余额 | [ ] |
| 05-011 | SharedWalletWalletInfoGetV30Api | 钱包信息 | [ ] |
| 05-012 | SharedWalletWalletOperationLogGetV30Api | 操作日志 | [ ] |
| 05-013 | SharedWalletWalletRelationGetV30Api | 关联关系 | [ ] |
| 05-014 | SharedWalletWatchRuleGetV30Api | 监控规则查询 | [ ] |
| 05-015 | SharedWalletWatchRuleSubmitV30Api | 监控规则提交 | [ ] |

## 06. 发票与对账

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 06-001 | CreateProjectInvoiceV2Api | 创建项目发票 | [ ] |
| 06-002 | CreateStatementInvoiceV2Api | 创建对账单发票 | [ ] |
| 06-003 | CreateStatementV2Api | 创建对账单 | [ ] |
| 06-004 | DownloadStatementV2Api | 下载对账单 | [ ] |
| 06-005 | DownloadStatementEsignFileV2Api | 下载电签文件 | [ ] |
| 06-006 | QueryInvoiceV2Api | 查询发票 | [ ] |
| 06-007 | QueryInvoiceSelfV2Api | 查询自开发票 | [ ] |
| 06-008 | QueryInvoiceDetailV2Api | 发票详情 | [ ] |
| 06-009 | QueryInvoiceDetailSelfV2Api | 自开发票详情 | [ ] |
| 06-010 | QueryInvoiceElectronicUrlV2Api | 电子发票URL | [ ] |
| 06-011 | QueryInvoiceElectronicUrlSelfV2Api | 自开电子发票URL | [ ] |
| 06-012 | QueryInvoiceTaxV2Api | 发票税务信息 | [ ] |
| 06-013 | QueryProjectV2Api | 查询项目 | [ ] |
| 06-014 | QueryProjectV30Api | 查询项目v3 | [ ] |
| 06-015 | QueryStatementV2Api | 查询对账单 | [ ] |
| 06-016 | QueryBookingBusinessEntityIdGetV2Api | 查询预订实体ID | [ ] |
| 06-017 | UploadStatementV2Api | 上传对账单 | [ ] |
| 06-018 | QueryRebateAccountingInfoV2Api | 返点结算信息 | [ ] |
| 06-019 | QueryRebateBalanceV2Api | 返点余额 | [ ] |

## 07. 代理商管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 07-001 | AgentInfoV2Api | 代理商信息 | [ ] |
| 07-002 | AgentAdvertiserSelectV2Api | 查询下属广告主 | [ ] |
| 07-003 | AgentAdvertiserInfoQueryV2Api | 下属广告主信息 | [ ] |
| 07-004 | AgentAdvertiserAssignV2Api | 分配广告主 | [ ] |
| 07-005 | AgentAdvertiserUnassignV2Api | 取消分配 | [ ] |
| 07-006 | AgentAdvertiserUpdateV2Api | 更新广告主 | [ ] |
| 07-007 | AgentAdvertiserCopyV2Api | 复制广告主 | [ ] |
| 07-008 | AgentChildAgentSelectV2Api | 二级代理商列表 | [ ] |
| 07-009 | AgentAdvAdvertiserUpdateSaleV2Api | 更新销售 | [ ] |
| 07-010 | AgentAdvBiddingListQueryV2Api | 竞价列表 | [ ] |
| 07-011 | AgentAdvBrandListQueryV2Api | 品牌列表 | [ ] |
| 07-012 | AgentAdvCostReportListQueryV2Api | 消耗报表 | [ ] |
| 07-013 | AgentAdvPerenniallyPunishV2Api | 长期处罚 | [ ] |
| 07-014 | AgentAdvPerenniallyPunishHistoryQueryV2Api | 处罚历史 | [ ] |
| 07-015 | AgentAdvTemporaryPunishV2Api | 临时处罚 | [ ] |
| 07-016 | AgentAdvRechargeRechargeRecordV2Api | 充值记录 | [ ] |
| 07-017 | AgentChargeVerifyV2Api | 充值验证 | [ ] |
| 07-018 | AgentCreditChargeSubmitV2Api | 授信充值 | [ ] |
| 07-019 | AgentCompanyBiddingListQueryV2Api | 公司竞价列表 | [ ] |
| 07-020 | AgentCompanyBrandListQueryV2Api | 公司品牌列表 | [ ] |
| 07-021 | AgentPrepayChargeGenerateRemittanceCodeV2Api | 汇款码 | [ ] |
| 07-022 | AgentQueryRiskPromotionListV2Api | 风险推广列表 | [ ] |
| 07-023 | AgentTransferTransactionRecordV2Api | 转账记录 | [ ] |

## 08. 客户中心/工作台

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 08-001 | CustomerCenterAccountListV30Api | 账户列表 | [ ] |
| 08-002 | CustomerCenterAccountOfflineListV30Api | 离线账户列表 | [ ] |
| 08-003 | CustomerCenterAdvertiserListV2Api | 广告主列表 | [ ] |
| 08-004 | CustomerCenterAdvertiserCopyV2Api | 复制广告主 | [ ] |
| 08-005 | CustomerCenterAdvertiserTransferableListV2Api | 可转移列表 | [ ] |
| 08-006 | BusinessPlatformCompanyAccountGetV30Api | 公司账户 | [ ] |
| 08-007 | BusinessPlatformCompanyInfoGetV30Api | 公司信息 | [ ] |
| 08-008 | BusinessPlatformPartnerOrganizationListV2Api | 合作组织列表 | [ ] |
| 08-009 | MajordomoAdvertiserSelectV2Api | 管家广告主列表 | [ ] |
| 08-010 | SubscribeAccountsAddV30Api | 订阅账户 | [ ] |
| 08-011 | SubscribeAccountsListV30Api | 订阅列表 | [ ] |
| 08-012 | SubscribeAccountsRemoveV30Api | 取消订阅 | [ ] |

## 09. 广告组 Campaign (v2)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 09-001 | CampaignCreateV2Api | 创建广告组 | [ ] |
| 09-002 | CampaignGetV2Api | 查询广告组 | [x] |
| 09-003 | CampaignUpdateV2Api | 更新广告组 | [ ] |
| 09-004 | CampaignUpdateStatusV2Api | 更新广告组状态 | [x] |

## 10. 广告计划 Ad (v2)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 10-001 | AdGetV2Api | 查询广告计划 | [x] |
| 10-002 | AdUpdateBidV2Api | 更新出价 | [x] |
| 10-003 | AdUpdateBudgetV2Api | 更新预算 | [x] |
| 10-004 | AdUdUpdateV2Api | 更新UD | [ ] |
| 10-005 | AdRejectReasonV2Api | 审核拒绝原因 | [x] |
| 10-006 | AdCostProtectStatusGetV2Api | 成本保护状态 | [x] |
| 10-007 | AdBillingUniqueIdGetV30Api | 计费唯一ID | [ ] |
| 10-008 | AdShopInfoUpdateV30Api | 更新店铺信息 | [ ] |
| 10-009 | AdConvertSignalV2Api | 转化信号 | [ ] |
| 10-010 | AdvConvertOleConvertV2Api | OLE转化 | [ ] |

## 11. 项目 Project (v3)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 11-001 | ProjectCreateV30Api | 创建项目 | [ ] |
| 11-002 | ProjectListV30Api | 项目列表 | [x] |
| 11-003 | ProjectUpdateV30Api | 更新项目 | [ ] |
| 11-004 | ProjectDeleteV30Api | 删除项目 | [x] |
| 11-005 | ProjectStatusUpdateV30Api | 更新状态 | [x] |
| 11-006 | ProjectBudgetUpdateV30Api | 更新预算 | [x] |
| 11-007 | ProjectCpaBidUpdateV30Api | 更新CPA出价 | [ ] |
| 11-008 | ProjectDeepCpaBidUpdateV30Api | 更新深度CPA出价 | [ ] |
| 11-009 | ProjectRoigoalUpdateV30Api | 更新ROI目标 | [ ] |
| 11-010 | ProjectNameUpdateV30Api | 更新名称 | [ ] |
| 11-011 | ProjectScheduleTimeUpdateV30Api | 更新排期 | [ ] |
| 11-012 | ProjectWeekScheduleUpdateV30Api | 更新周排期 | [ ] |
| 11-013 | ProjectCostProtectStatusGetV30Api | 成本保护状态 | [x] |

## 12. 广告单元 Promotion (v3)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 12-001 | PromotionCreateV30Api | 创建单元 | [ ] |
| 12-002 | PromotionListV30Api | 单元列表 | [x] |
| 12-003 | PromotionUpdateV30Api | 更新单元 | [ ] |
| 12-004 | PromotionDeleteV30Api | 删除单元 | [x] |
| 12-005 | PromotionStatusUpdateV30Api | 更新状态 | [x] |
| 12-006 | PromotionBidUpdateV30Api | 更新出价 | [x] |
| 12-007 | PromotionBudgetUpdateV30Api | 更新预算 | [x] |
| 12-008 | PromotionDeepbidUpdateV30Api | 更新深度出价 | [ ] |
| 12-009 | PromotionNameUpdateV30Api | 更新名称 | [ ] |
| 12-010 | PromotionScheduleTimeUpdateV30Api | 更新排期 | [ ] |
| 12-011 | PromotionShopInfoUpdateV30Api | 更新店铺信息 | [ ] |
| 12-012 | PromotionMaterialDeleteV30Api | 删除素材 | [ ] |
| 12-013 | PromotionRejectReasonGetV30Api | 审核拒绝原因 | [x] |
| 12-014 | PromotionCostProtectStatusGetV30Api | 成本保护状态 | [x] |
| 12-015 | PromotionAidGetV30Api | AID查询 | [ ] |
| 12-016 | PromotionAutoGenerateConfigCreateV30Api | 自动生成配置 | [ ] |
| 12-017 | PromotionAutoGenerateConfigGetV30Api | 查询自动生成配置 | [ ] |
| 12-018 | PromotionEasyUpdateV30Api | 简易更新 | [ ] |
| 12-019 | PromotionEasyKeepDeliverySwitchUpdateV30Api | 保投开关 | [ ] |
| 12-020 | PromotionNewcustomerCreateV30Api | 新客创建 | [ ] |
| 12-021 | PromotionNewcustomerTypeGetV30Api | 新客类型查询 | [ ] |

## 13. 预算组

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 13-001 | BudgetGroupCreateV30Api | 创建预算组 | [ ] |
| 13-002 | BudgetGroupListV30Api | 预算组列表 | [ ] |
| 13-003 | BudgetGroupUpdateV30Api | 更新预算组 | [ ] |
| 13-004 | BudgetGroupDeleteV30Api | 删除预算组 | [ ] |

## 14. 全域推广 UniProject

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 14-001 | UniProjectCreateV30Api | 创建全域项目 | [ ] |
| 14-002 | UniProjectListGetV30Api | 全域项目列表 | [ ] |
| 14-003 | UniProjectUpdateV30Api | 更新全域项目 | [ ] |
| 14-004 | UniProjectMonetizationModeUpdateV30Api | 更新变现模式 | [ ] |
| 14-005 | UniProjectAwemeAuthorizedGetV30Api | 授权抖音号 | [ ] |

## 15. 关键词

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 15-001 | KeywordCreateV2V2Api | 创建关键词v2 | [ ] |
| 15-002 | KeywordCreateV30Api | 创建关键词v3 | [ ] |
| 15-003 | KeywordGetV2Api | 查询关键词 | [ ] |
| 15-004 | KeywordListV30Api | 关键词列表 | [ ] |
| 15-005 | KeywordUpdateV2V2Api | 更新关键词v2 | [ ] |
| 15-006 | KeywordUpdateV30Api | 更新关键词v3 | [ ] |
| 15-007 | KeywordDeleteV2V2Api | 删除关键词v2 | [ ] |
| 15-008 | KeywordDeleteV30Api | 删除关键词v3 | [ ] |
| 15-009 | KeywordFeedadsSuggestV2Api | 信息流关键词建议 | [ ] |

## 16. 创意 Creative

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 16-001 | CreativeGetV2Api | 查询创意 | [x] |
| 16-002 | CreativeDetailGetV30Api | 创意详情 | [x] |
| 16-003 | CreativeCustomCreativeCreateV2Api | 创建自定义创意 | [ ] |
| 16-004 | CreativeCustomCreativeUpdateV2Api | 更新自定义创意 | [ ] |
| 16-005 | CreativeProceduralCreativeCreateV2Api | 创建程序化创意 | [ ] |
| 16-006 | CreativeProceduralCreativeUpdateV2Api | 更新程序化创意 | [ ] |
| 16-007 | CreativeRejectReasonV2Api | 创意审核拒绝原因 | [x] |
| 16-008 | CreativeStrategyListV2Api | 创意策略列表 | [ ] |
| 16-009 | AssetsCreativeComponentCreateV2Api | 创建创意组件 | [ ] |
| 16-010 | AssetsCreativeComponentGetV2Api | 查询创意组件 | [ ] |
| 16-011 | AssetsCreativeComponentUpdateV2Api | 更新创意组件 | [ ] |

## 17. 文件/素材管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 17-001 | FileImageAdV2Api | 上传广告图片 | [ ] |
| 17-002 | FileImageAdGetV2Api | 查询广告图片 | [ ] |
| 17-003 | FileImageGetV2Api | 查询图片 | [ ] |
| 17-004 | FileImageAdvertiserV2Api | 上传广告主图片 | [ ] |
| 17-005 | FileImageDeleteV30Api | 删除图片 | [ ] |
| 17-006 | FileVideoAdV2Api | 上传广告视频 | [ ] |
| 17-007 | FileVideoAdGetV2Api | 查询广告视频 | [ ] |
| 17-008 | FileVideoGetV2Api | 查询视频 | [ ] |
| 17-009 | FileVideoDeleteV2Api | 删除视频 | [ ] |
| 17-010 | FileVideoUpdateV2Api | 更新视频 | [ ] |
| 17-011 | FileVideoPauseV2Api | 暂停视频 | [ ] |
| 17-012 | FileVideoAwemeGetV2Api | 查询抖音视频 | [ ] |
| 17-013 | FileVideoAgentV2Api | 代理商上传视频 | [ ] |
| 17-014 | FileVideoAgentGetV2Api | 代理商查询视频 | [ ] |
| 17-015 | FileVideoEfficiencyGetV2Api | 视频效率分析 | [ ] |
| 17-016 | FileVideoMaterialClearTaskCreateV2Api | 创建素材清理任务 | [ ] |
| 17-017 | FileVideoMaterialClearTaskGetV2Api | 查询清理任务 | [ ] |
| 17-018 | FileVideoMaterialClearTaskResultGetV2Api | 清理任务结果 | [ ] |
| 17-019 | FileVideoUploadTaskListV2Api | 上传任务列表 | [ ] |
| 17-020 | FileUploadTaskCreateV2Api | 创建上传任务 | [ ] |
| 17-021 | FileAudioAdV2Api | 上传广告音频 | [ ] |
| 17-022 | FileAudioGetV2Api | 查询音频 | [ ] |
| 17-023 | FileAutoGenerateSourceGetV2Api | 自动生成素材来源 | [ ] |
| 17-024 | FileCarouselAwemeGetV30Api | 抖音轮播图 | [ ] |
| 17-025 | FileEbpVideoGetV30Api | EBP视频查询 | [ ] |
| 17-026 | FileMaterialAttributesListV2Api | 素材属性列表 | [ ] |
| 17-027 | FileMaterialBindV2Api | 绑定素材 | [ ] |
| 17-028 | FileMaterialDetailV2Api | 素材详情 | [ ] |
| 17-029 | FileMaterialListV2Api | 素材列表 | [ ] |
| 17-030 | FilePlayableCreateV30Api | 创建可玩广告 | [ ] |
| 17-031 | FilePlayableListV30Api | 可玩广告列表 | [ ] |
| 17-032 | FilePreauditGetV30Api | 预审查询 | [ ] |
| 17-033 | FilePreauditSubmitV30Api | 提交预审 | [ ] |
| 17-034 | FileQualityGetV30Api | 质量查询 | [ ] |
| 17-035 | FileQualitySubmitV30Api | 提交质量检测 | [ ] |
| 17-036 | MaterialStatusUpdateV30Api | 素材状态更新 | [ ] |
| 17-037 | FileRebateCommonDownloadCreateTaskV2Api | 返点下载任务 | [ ] |
| 17-038 | FileRebateCommonDownloadDownloadFileV2Api | 返点下载文件 | [ ] |
| 17-039 | FileRebateCommonDownloadGetDownloadTaskListV2Api | 返点下载列表 | [ ] |
| 17-040 | FileRebateMaterialDownloadCreateTaskV2Api | 返点素材下载 | [ ] |
| 17-041 | FileRebateMaterialDownloadDownloadFileV2Api | 返点素材文件 | [ ] |
| 17-042 | FileRebateMaterialDownloadGetDownloadTaskListV2Api | 返点素材列表 | [ ] |
| 17-043 | FileRebateRebateDownloadCreateTaskV2Api | 返点下载任务 | [ ] |

## 18. AI创意 AIC

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 18-001 | AicElementUploadV30Api | 上传元素 | [ ] |
| 18-002 | AicElementGetV30Api | 查询元素 | [ ] |
| 18-003 | AicElementUpdateV30Api | 更新元素 | [ ] |
| 18-004 | AicElementDeleteV30Api | 删除元素 | [ ] |
| 18-005 | AicImageMixcutCreateV30Api | 图片混剪 | [ ] |
| 18-006 | AicVideoMixcutCreateV30Api | 视频混剪 | [ ] |
| 18-007 | AicMixcutTaskSaveV30Api | 保存混剪任务 | [ ] |
| 18-008 | AicMixcutTaskResultGetV30Api | 混剪结果 | [ ] |
| 18-009 | AicMaterialGetV30Api | AI素材查询 | [ ] |
| 18-010 | AicMaterialPushV30Api | AI素材推送 | [ ] |

## 19. 轮播图

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 19-001 | CarouselCreateV2Api | 创建轮播图 | [ ] |
| 19-002 | CarouselListV2Api | 轮播图列表 | [ ] |
| 19-003 | CarouselUpdateV2Api | 更新轮播图 | [ ] |
| 19-004 | CarouselDeleteV2Api | 删除轮播图 | [ ] |
| 19-005 | CarouselAdGetV2Api | 轮播图广告查询 | [ ] |

## 20. 数据报表

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 20-001 | ReportAdvertiserGetV2Api | 广告主报表v2 | [ ] (40118 未授权) |
| 20-002 | ReportCampaignGetV2Api | 广告组报表v2 | [ ] (40118 未授权) |
| 20-003 | ReportAdGetV2Api | 广告计划报表v2 | [ ] (40118 未授权) |
| 20-004 | ReportCreativeGetV2Api | 创意报表v2 | [ ] |
| 20-005 | ReportCustomGetV30Api | 自定义报表v3 | [x] |
| 20-006 | ReportCustomCreativeGetV30Api | 自定义创意报表 | [ ] |
| 20-007 | ReportCustomConfigGetV30Api | 报表配置查询 | [x] |
| 20-008 | ReportCustomAsyncTaskCreateV30Api | 创建异步报表任务 | [ ] |
| 20-009 | ReportCustomAsyncTaskGetV30Api | 查询异步任务状态 | [ ] |
| 20-010 | ReportCustomAsyncTaskDownloadV30Api | 下载异步报表 | [ ] |
| 20-011 | ReportVideoFrameGetV2Api | 视频帧报表 | [ ] |
| 20-012 | ReportSitePageV2Api | 落地页报表 | [ ] |
| 20-013 | ReportAgentGetV2V2Api | 代理商报表 | [ ] |
| 20-014 | ReportRubeexGetV2Api | Rubeex报表 | [ ] |

## 21. 受众分析报表

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 21-001 | ReportAudienceAgeV2Api | 年龄分布 | [ ] |
| 21-002 | ReportAudienceGenderV2Api | 性别分布 | [ ] |
| 21-003 | ReportAudienceProvinceV2Api | 省份分布 | [ ] |
| 21-004 | ReportAudienceCityV2Api | 城市分布 | [ ] |
| 21-005 | ReportAudienceInterestActionListV2Api | 兴趣行为分布 | [ ] |
| 21-006 | ReportAudienceAwemeListV2Api | 达人分布 | [ ] |

## 22. 品牌报表

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 22-001 | ReportBrandAdvertiserGetV30Api | 品牌广告主报表 | [ ] |
| 22-002 | ReportBrandCampaignGetV30Api | 品牌广告组报表 | [ ] |
| 22-003 | ReportBrandAdGetV30Api | 品牌广告报表 | [ ] |
| 22-004 | ReportBrandCreativeGetV30Api | 品牌创意报表 | [ ] |
| 22-005 | ReportBrandDataV30Api | 品牌数据 | [ ] |
| 22-006 | ReportBrandAgentDataV30Api | 品牌代理数据 | [ ] |

## 23. 直播/商品/其他报表

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 23-001 | ReportLiveRoomAnalysisGetV30Api | 直播间分析 | [ ] |
| 23-002 | ReportReportLiveRoomAudiencePortraitGetV30Api | 直播受众画像 | [ ] |
| 23-003 | ReportProductDailyAsyncTaskCreateV30Api | 商品日报任务 | [ ] |
| 23-004 | ReportProductHourlyAsyncTaskCreateV30Api | 商品时报任务 | [ ] |
| 23-005 | ReportProductAsyncTaskGetV30Api | 商品任务状态 | [ ] |
| 23-006 | ReportProductAsyncTaskDownloadV30Api | 下载商品报表 | [ ] |
| 23-007 | ReportJointGrowthGetV30Api | 联合增长报表 | [ ] |
| 23-008 | ReportJointGrowthCusGetV30Api | 联合增长客户报表 | [ ] |
| 23-009 | ReportStardeliveryTaskDataGetV30Api | 星推任务数据 | [ ] |
| 23-010 | ReportStardeliveryTaskVideoDataGetV30Api | 星推视频数据 | [ ] |
| 23-011 | ReportBusinessPlatformStardeliveryTaskVideoDataGetV30Api | 平台星推视频 | [ ] |

## 24. RTA报表

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 24-001 | ReportRtaGetV2Api | RTA报表 | [ ] |
| 24-002 | ReportRtaExpGetV2Api | RTA实验报表 | [ ] |
| 24-003 | ReportRtaCusExpGetV2Api | RTA客户实验报表 | [ ] |
| 24-004 | ReportRtaExpLocalDailyGetV30Api | RTA本地日报 | [ ] |
| 24-005 | ReportRtaExpLocalHourlyGetV30Api | RTA本地时报 | [ ] |

## 25. 人群包 DMP

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 25-001 | DmpCustomAudienceSelectV2Api | 人群包列表 | [ ] |
| 25-002 | DmpCustomAudienceReadV2Api | 人群包详情 | [ ] |
| 25-003 | DmpCustomAudiencePublishV2Api | 发布人群包 | [ ] |
| 25-004 | DmpCustomAudiencePushV2V2Api | 推送人群包 | [ ] |
| 25-005 | DmpCustomAudienceCopyV2Api | 复制人群包 | [ ] |
| 25-006 | DmpCustomAudienceDeleteV2Api | 删除人群包 | [ ] |
| 25-007 | DmpDataSourceCreateV2Api | 创建数据源 | [ ] |
| 25-008 | DmpDataSourceReadV2Api | 查询数据源 | [ ] |
| 25-009 | DmpDataSourceUpdateV2Api | 更新数据源 | [ ] |
| 25-010 | DmpDataSourceFileUploadV2Api | 上传数据源文件 | [ ] |
| 25-011 | DmpBrandGetV2Api | 品牌DMP查询 | [ ] |
| 25-012 | AudiencePackageCreateV2Api | 创建定向包 | [ ] |
| 25-013 | AudiencePackageGetV30Api | 查询定向包 | [ ] |
| 25-014 | AudiencePackageUpdateV2Api | 更新定向包 | [ ] |
| 25-015 | AudiencePackageDeleteV2Api | 删除定向包 | [ ] |
| 25-016 | AudiencePackageBindinfoGetV30Api | 定向包绑定信息 | [ ] |

## 26. 转化追踪 EventManager

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 26-001 | EventManagerAssetsCreateV2Api | 创建资产 | [ ] |
| 26-002 | EventManagerAbnormalAssetsGetV30Api | 异常资产查询 | [ ] |
| 26-003 | EventManagerAvailableEventsGetV2Api | 可用事件查询 | [ ] |
| 26-004 | EventManagerEventsCreateV2Api | 创建事件 | [ ] |
| 26-005 | EventManagerEventConfigsGetV2Api | 事件配置查询 | [ ] |
| 26-006 | EventManagerDeepBidTypeGetV30Api | 深度出价类型 | [ ] |
| 26-007 | EventManagerOptimizedGoalGetV2V30Api | 优化目标查询 | [ ] |
| 26-008 | EventManagerTrackUrlCreateV2Api | 创建监测链接 | [ ] |
| 26-009 | EventManagerTrackUrlGetV2Api | 查询监测链接 | [ ] |
| 26-010 | EventManagerTrackUrlUpdateV2Api | 更新监测链接 | [ ] |
| 26-011 | EventManagerAuthEnableV2Api | 启用认证 | [ ] |
| 26-012 | EventManagerAuthDisableV2Api | 禁用认证 | [ ] |
| 26-013 | EventManagerAuthGetAuthStatusV2Api | 认证状态 | [ ] |
| 26-014 | EventManagerAuthAddPublicKeyV2Api | 添加公钥 | [ ] |
| 26-015 | EventManagerAuthDelPublicKeyV2Api | 删除公钥 | [ ] |
| 26-016 | EventManagerAuthGetAllPublicKeysV2Api | 所有公钥 | [ ] |
| 26-017 | EventManagerAuthGetPublicKeyV2Api | 查询公钥 | [ ] |
| 26-018 | EventManagerShareV30Api | 共享资产 | [ ] |
| 26-019 | EventManagerShareGetV30Api | 查询共享 | [ ] |
| 26-020 | EventManagerShareCancelV30Api | 取消共享 | [ ] |
| 26-021 | AnalyticsAttributionV30Api | 归因分析 | [ ] |

## 27. 线索管理 Clue

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 27-001 | ClueFormCreateV2Api | 创建表单 | [ ] |
| 27-002 | ClueFormListV2Api | 表单列表 | [ ] |
| 27-003 | ClueFormDetailV2Api | 表单详情 | [ ] |
| 27-004 | ClueFormUpdateV2Api | 更新表单 | [ ] |
| 27-005 | ClueFormDeleteV2Api | 删除表单 | [ ] |
| 27-006 | ClueCaCreateV2Api | 创建CA | [ ] |
| 27-007 | ClueCaUpdateV2Api | 更新CA | [ ] |
| 27-008 | ClueCaInterfaceCreateV2Api | 创建CA接口 | [ ] |
| 27-009 | ClueCaInterfaceUpdateV2Api | 更新CA接口 | [ ] |
| 27-010 | ClueSmartphoneCreateV2Api | 创建智能电话 | [ ] |
| 27-011 | ClueSmartphoneGetV2Api | 查询智能电话 | [ ] |
| 27-012 | ClueSmartphoneRecordV2Api | 电话记录 | [ ] |
| 27-013 | ClueSmartphoneDeleteV2Api | 删除智能电话 | [ ] |
| 27-014 | ClueCouponCreateV2Api | 创建优惠券 | [ ] |
| 27-015 | ClueCouponGetV2Api | 查询优惠券 | [ ] |
| 27-016 | ClueCouponDetailV2Api | 优惠券详情 | [ ] |
| 27-017 | ClueCouponUpdateV2Api | 更新优惠券 | [ ] |
| 27-018 | ClueCouponCodeGetV2Api | 优惠券码查询 | [ ] |
| 27-019 | ClueCouponCodeConsumeV2Api | 核销优惠券 | [ ] |
| 27-020 | ClueCouponEmployeeCreateV2Api | 添加员工 | [ ] |
| 27-021 | ClueCouponEmployeeGetV2Api | 查询员工 | [ ] |
| 27-022 | ClueCouponEmployeeDeleteV2Api | 删除员工 | [ ] |
| 27-023 | ClueWechatDataGetV2Api | 微信线索数据 | [ ] |
| 27-024 | ClueWechatInstanceListV2Api | 微信实例列表 | [ ] |
| 27-025 | ClueWechatInstanceDetailV2Api | 微信实例详情 | [ ] |
| 27-026 | ClueWechatInstanceUpdateV2Api | 更新微信实例 | [ ] |
| 27-027 | ClueWechatPoolListV2Api | 微信池列表 | [ ] |

## 28. 落地页建站

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 28-001 | ToolsSiteCreateV2Api | 创建站点 | [ ] |
| 28-002 | ToolsSiteGetV2Api | 查询站点 | [ ] |
| 28-003 | ToolsSiteReadV2Api | 读取站点 | [ ] |
| 28-004 | ToolsSiteUpdateV2Api | 更新站点 | [ ] |
| 28-005 | ToolsSiteUpdateStatusV2Api | 更新站点状态 | [ ] |
| 28-006 | ToolsSiteCopyV2Api | 复制站点 | [ ] |
| 28-007 | ToolsSitePreviewV2Api | 预览站点 | [ ] |
| 28-008 | ToolsSiteHandselV2Api | 站点手选 | [ ] |
| 28-009 | ToolsSiteFormsListV2Api | 站点表单列表 | [ ] |
| 28-010 | ToolsSiteTemplateCreateV2Api | 创建模板 | [ ] |
| 28-011 | ToolsSiteTemplateGetV2Api | 查询模板 | [ ] |
| 28-012 | ToolsSiteTemplatePicUrlGetV2Api | 模板图片URL | [ ] |
| 28-013 | ToolsSiteTemplatePreviewV2Api | 预览模板 | [ ] |
| 28-014 | ToolsSiteTemplateSiteCreateV2Api | 从模板建站 | [ ] |
| 28-015 | ToolsAipThirdSiteCreateV2Api | 创建第三方站点 | [ ] |
| 28-016 | ToolsAipThirdSiteGetV2Api | 查询第三方站点 | [ ] |
| 28-017 | ToolsAipThirdSiteUpdateV2Api | 更新第三方站点 | [ ] |
| 28-018 | ToolsThirdSiteDeleteV2Api | 删除第三方站点 | [ ] |
| 28-019 | ToolsThirdSiteGetV2Api | 查询第三方站点v2 | [ ] |
| 28-020 | ToolsThirdSitePreviewV2Api | 预览第三方站点 | [ ] |
| 28-021 | ToolsThirdSiteUpdateV2Api | 更新第三方站点v2 | [ ] |
| 28-022 | ToolsOrangeSiteGetV30Api | 橙子建站查询 | [ ] |
| 28-023 | ToolsLandingGroupCreateV2Api | 创建落地页组 | [ ] |
| 28-024 | ToolsLandingGroupGetV2Api | 查询落地页组 | [ ] |
| 28-025 | ToolsLandingGroupUpdateV2Api | 更新落地页组 | [ ] |
| 28-026 | ToolsLandingGroupSiteOptStatusUpdateV2Api | 落地页优化状态 | [ ] |

## 29. 线索工具 (青鸟)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 29-001 | ToolsClueGetV2Api | 查询线索 | [ ] |
| 29-002 | ToolsClueInfoGetV2Api | 线索信息 | [ ] |
| 29-003 | ToolsClueInfoUpdateV2Api | 更新线索 | [ ] |
| 29-004 | ToolsClueFormGetV2Api | 线索表单 | [ ] |
| 29-005 | ToolsClueFormDetailV2Api | 线索表单详情 | [ ] |
| 29-006 | ToolsClueClueOverviewQueryV2Api | 线索总览 | [ ] |
| 29-007 | ToolsClueCallbackV2Api | 线索回传 | [ ] |
| 29-008 | ToolsClueExtInfoCallbackV2Api | 线索扩展信息回传 | [ ] |
| 29-009 | ToolsClueLifeCallbackV2Api | 线索生命周期回传 | [ ] |
| 29-010 | ToolsClueLifeGetV2Api | 线索生命周期查询 | [ ] |
| 29-011 | ToolsCluePrivateMessageCallbackV2Api | 私信线索回传 | [ ] |
| 29-012 | ToolsClueCallCreateV2Api | 创建呼叫 | [ ] |
| 29-013 | ToolsClueBridgeCallCreateV2Api | 桥接呼叫 | [ ] |
| 29-014 | ToolsClueCallVirtualNumberGetV2Api | 虚拟号码 | [ ] |
| 29-015 | ToolsClueCallVirtualNumberRefundDetailGetV2Api | 退款详情 | [ ] |
| 29-016 | ToolsClueContactLogListV2Api | 联系日志列表 | [ ] |
| 29-017 | ToolsClueContactLogOverviewQueryV2Api | 联系日志总览 | [ ] |
| 29-018 | ToolsClueContactLogRecordUrlGetV2Api | 通话录音URL | [ ] |
| 29-019 | ToolsClueLiteContactGetV2Api | 轻量联系查询 | [ ] |
| 29-020 | ToolsClueLiteContactRecordV2Api | 轻量联系记录 | [ ] |
| 29-021 | ToolsClueSmartPhoneGetV2Api | 智能电话查询 | [ ] |
| 29-022 | ToolsClueRobotScriptQueryV2Api | 机器人脚本查询 | [ ] |
| 29-023 | ToolsClueRobotTaskCreateV2Api | 创建机器人任务 | [ ] |
| 29-024 | ToolsClueRobotTaskCancelV2Api | 取消机器人任务 | [ ] |
| 29-025 | ToolsClueRefundInfoQueryV2Api | 退款信息查询 | [ ] |
| 29-026 | ToolsClueRefundDetailGetV2Api | 退款详情 | [ ] |
| 29-027 | ToolsClueRefundReportGetV2Api | 退款报表 | [ ] |
| 29-028 | ToolsClueRefundViewGetV2Api | 退款视图 | [ ] |
| 29-029 | ToolsClueWebrtcCreateV2V30Api | WebRTC创建 | [ ] |
| 29-030 | ToolsClueWebrtcTokenGetV2Api | WebRTC Token | [ ] |
| 29-031 | ToolsClueWebrtcTokenGetV2V30Api | WebRTC Token v3 | [ ] |

## 30. 抖音号管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 30-001 | ToolsAwemeAuthV2Api | 抖音号授权 | [ ] |
| 30-002 | ToolsAwemeAuthListV2Api | 授权列表 | [ ] |
| 30-003 | ToolsAwemeAuthCancelV2Api | 取消授权 | [ ] |
| 30-004 | ToolsAwemeAuthRenewalV2Api | 续期授权 | [ ] |
| 30-005 | ToolsAwemeAuthAuthShareAdShareV2Api | 分享授权 | [ ] |
| 30-006 | ToolsAwemeAuthorInfoGetV2Api | 达人信息 | [ ] |
| 30-007 | ToolsAwemeInfoSearchV2Api | 搜索抖音号 | [ ] |
| 30-008 | ToolsAwemeCategoryTopAuthorGetV2Api | 类目TOP达人 | [ ] |
| 30-009 | ToolsAwemeMultiLevelCategoryGetV2Api | 多级类目 | [ ] |
| 30-010 | ToolsAwemeSimilarAuthorSearchV2Api | 相似达人搜索 | [ ] |
| 30-011 | ToolsAwemeBannedCreateV30Api | 创建屏蔽 | [ ] |
| 30-012 | ToolsAwemeBannedListV30Api | 屏蔽列表 | [ ] |
| 30-013 | ToolsAwemeBannedDeleteV30Api | 删除屏蔽 | [ ] |
| 30-014 | ToolsLiveAuthorizeListV2Api | 直播授权列表 | [ ] |

## 31. 评论管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 31-001 | ToolsCommentGetV30Api | 查询评论 | [ ] |
| 31-002 | ToolsCommentReplyV30Api | 回复评论 | [ ] |
| 31-003 | ToolsCommentReplyGetV30Api | 查询回复 | [ ] |
| 31-004 | ToolsCommentHideV30Api | 隐藏评论 | [ ] |
| 31-005 | ToolsCommentStickOnTopV30Api | 置顶评论 | [ ] |
| 31-006 | ToolsCommentMetricsGetV30Api | 评论指标 | [ ] |
| 31-007 | ToolsCommentMid2itemIdV30Api | MID转ItemID | [ ] |
| 31-008 | ToolsCommentTermsBannedAddV30Api | 添加屏蔽词 | [ ] |
| 31-009 | ToolsCommentTermsBannedGetV30Api | 查询屏蔽词 | [ ] |
| 31-010 | ToolsCommentTermsBannedUpdateV30Api | 更新屏蔽词 | [ ] |
| 31-011 | ToolsCommentTermsBannedDeleteV30Api | 删除屏蔽词 | [ ] |

## 32. 工具-出价/预估/诊断

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 32-001 | ToolsBidSuggestV2Api | 出价建议v2 | [ ] |
| 32-002 | ToolsBidsSuggestV30Api | 出价建议v3 | [ ] |
| 32-003 | ToolsNoBidSuggestBidV2Api | 无出价建议 | [ ] |
| 32-004 | ToolsEstimateAudienceV2Api | 受众规模预估 | [ ] |
| 32-005 | ToolsEstimatedPriceGetV2Api | 预估价格 | [ ] |
| 32-006 | ToolsSuggestBudgetGetV30Api | 预算建议 | [ ] |
| 32-007 | ToolsSearchBidRatioGetV2Api | 搜索出价比 | [ ] |
| 32-008 | ToolsAdvertiserDiagnosisSuggestionGetV30Api | 广告主诊断 | [ ] |
| 32-009 | ToolsAdvertiserDiagnosisSuggestionAcceptUpdateV30Api | 采纳诊断 | [ ] |
| 32-010 | ToolsPromotionDiagnosisSuggestionGetV30Api | 单元诊断 | [ ] |
| 32-011 | ToolsPromotionDiagnosisSuggestionAcceptV30Api | 采纳单元诊断 | [ ] |
| 32-012 | ToolsDiagnosisSuggestionGetV30Api | 诊断建议 | [ ] |
| 32-013 | DiagnosisTaskAdvCreateV2Api | 创建诊断任务 | [ ] |
| 32-014 | DiagnosisTaskAdvGetV2Api | 查询诊断任务 | [ ] |
| 32-015 | DiagnosisTaskAdvListV2Api | 诊断任务列表 | [ ] |
| 32-016 | DiagnosisTaskAgentCreateV2Api | 代理商诊断任务 | [ ] |
| 32-017 | DiagnosisTaskAgentGetV2Api | 查询代理商诊断 | [ ] |
| 32-018 | DiagnosisTaskAgentListV2Api | 代理商诊断列表 | [ ] |

## 33. 工具-地域/行业/创意词

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 33-001 | ToolsRegionGetV2Api | 地域列表 | [ ] |
| 33-002 | ToolsCountryInfoV2Api | 国家信息 | [ ] |
| 33-003 | ToolsIndustryGetV2Api | 行业列表 | [ ] |
| 33-004 | ToolsCreativeWordSelectV2Api | 创意词查询 | [ ] |
| 33-005 | ToolsInterestActionActionKeywordV2Api | 行为关键词 | [ ] |
| 33-006 | ToolsInterestActionInterestKeywordV2Api | 兴趣关键词 | [ ] |
| 33-007 | ToolsInterestActionKeywordSuggestV2Api | 关键词建议 | [ ] |
| 33-008 | ToolsInterestActionId2wordV2Api | ID转词 | [ ] |
| 33-009 | SuggWordsV30Api | 推荐词 | [ ] |

## 34. 一键起量/素材起量

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 34-001 | ToolsTaskRaiseCreateV2Api | 创建起量任务 | [ ] |
| 34-002 | ToolsTaskRaiseGetV2Api | 查询起量任务 | [ ] |
| 34-003 | ToolsTaskRaiseDataGetV2Api | 起量数据 | [ ] |
| 34-004 | ToolsTaskRaiseOptimizationIdsGetV2Api | 优化ID列表 | [ ] |
| 34-005 | ToolsTaskRaiseStatusStopV2Api | 停止起量 | [ ] |
| 34-006 | ToolsAdRaiseStatusGetV2Api | 广告起量状态 | [ ] |
| 34-007 | ToolsPromotionRaiseSetV30Api | 单元起量设置 | [ ] |
| 34-008 | ToolsPromotionRaiseStatusGetV30Api | 单元起量状态 | [ ] |
| 34-009 | ToolsPromotionRaiseStatusCurrentIdsGetV30Api | 当前起量ID | [ ] |
| 34-010 | ToolsPromotionRaiseStopV30Api | 停止单元起量 | [ ] |
| 34-011 | ToolsPromotionRaiseVersionGetV30Api | 起量版本 | [ ] |
| 34-012 | ToolsMaterialRaiseCreateV30Api | 创建素材起量 | [ ] |
| 34-013 | ToolsMaterialRaiseGetV30Api | 查询素材起量 | [ ] |
| 34-014 | ToolsMaterialRaiseMaterialIdsGetV30Api | 素材起量ID | [ ] |
| 34-015 | ToolsMaterialRaiseRecordsGetV30Api | 素材起量记录 | [ ] |
| 34-016 | ToolsMaterialRaiseStatusGetV30Api | 素材起量状态 | [ ] |
| 34-017 | ToolsMaterialRaiseStopV30Api | 停止素材起量 | [ ] |

## 35. 搜索/关键词工具

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 35-001 | ToolsBlueFlowKeywordListV30Api | 蓝海关键词 | [ ] |
| 35-002 | ToolsBlueFlowPackageListV30Api | 蓝海流量包 | [ ] |
| 35-003 | ToolsKeywordsBidRatioCreateV30Api | 创建出价比 | [ ] |
| 35-004 | ToolsKeywordsBidRatioGetV30Api | 查询出价比 | [ ] |
| 35-005 | ToolsKeywordsBidRatioUpdateV30Api | 更新出价比 | [ ] |
| 35-006 | ToolsKeywordsBidRatioDeleteV30Api | 删除出价比 | [ ] |
| 35-007 | ToolsKeywordsProjectInfoGetV30Api | 关键词项目信息 | [ ] |
| 35-008 | ToolsPrivativeWordGetV2Api | 否定词查询 | [ ] |
| 35-009 | ToolsPrivativeWordAdAddV2Api | 广告否定词 | [ ] |
| 35-010 | ToolsPrivativeWordAdUpdateV2Api | 更新广告否定词 | [ ] |
| 35-011 | ToolsPrivativeWordCampaignAddV2Api | 广告组否定词 | [ ] |
| 35-012 | ToolsPrivativeWordCampaignUpdateV2Api | 更新广告组否定词 | [ ] |
| 35-013 | ToolsPrivativeWordBatchGetV30Api | 批量否定词 | [ ] |
| 35-014 | ToolsPrivativeWordProjectAddV30Api | 项目否定词 | [ ] |
| 35-015 | ToolsPrivativeWordProjectUpdateV30Api | 更新项目否定词 | [ ] |
| 35-016 | ToolsPrivativeWordPromotionAddV30Api | 单元否定词 | [ ] |
| 35-017 | ToolsPrivativeWordPromotionUpdateV30Api | 更新单元否定词 | [ ] |
| 35-018 | ToolsJointBidCreateV30Api | 联合出价 | [ ] |
| 35-019 | ToolsJointBidGetV30Api | 查询联合出价 | [ ] |
| 35-020 | ToolsJointBidUpdateV30Api | 更新联合出价 | [ ] |

## 36. APP/小程序/小游戏管理

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 36-001 | ToolsAppManagementAppGetV2Api | 应用查询 | [ ] |
| 36-002 | ToolsAppManagementAndroidAppListV2Api | 安卓应用列表 | [ ] |
| 36-003 | ToolsAppManagementHarmonyAppListV2Api | 鸿蒙应用列表 | [ ] |
| 36-004 | ToolsAppIosListV2Api | iOS应用列表 | [ ] |
| 36-005 | ToolsAppManagementAndroidBasicPackageGetV2Api | 安卓包查询 | [ ] |
| 36-006 | ToolsAppManagementAndroidBasicPackagePublishV2Api | 发布安卓包 | [ ] |
| 36-007 | ToolsAppManagementAndroidBasicPackageUpdateV2Api | 更新安卓包 | [ ] |
| 36-008 | ToolsAppManagementExtendPackageCreateV2Api | 创建扩展包 | [ ] |
| 36-009 | ToolsAppManagementExtendPackageCreateV2V2Api | 创建扩展包v2 | [ ] |
| 36-010 | ToolsAppManagementExtendPackageListV2Api | 扩展包列表 | [ ] |
| 36-011 | ToolsAppManagementExtendPackageListV2V2Api | 扩展包列表v2 | [ ] |
| 36-012 | ToolsAppManagementExtendPackageUpdateV2Api | 更新扩展包 | [ ] |
| 36-013 | ToolsAppManagementExtendPackageUpdateV2V2Api | 更新扩展包v2 | [ ] |
| 36-014 | ToolsAppManagementBookingGetV2Api | 预约查询 | [ ] |
| 36-015 | ToolsAppManagementBookingRecordsGetV2Api | 预约记录 | [ ] |
| 36-016 | ToolsAppManagementBpShareV2Api | BP分享 | [ ] |
| 36-017 | ToolsAppManagementBpShareCancelV2Api | 取消分享 | [ ] |
| 36-018 | ToolsAppManagementShareAccountListV2Api | 分享账户列表 | [ ] |
| 36-019 | ToolsAppManagementUpdateAuthorizationV2Api | 更新授权 | [ ] |
| 36-020 | ToolsAppManagementIndustryInfoListV2Api | 行业信息 | [ ] |
| 36-021 | ToolsAppManagementUploadTaskCreateV2Api | 创建上传任务 | [ ] |
| 36-022 | ToolsAppManagementUploadTaskListV2Api | 上传任务列表 | [ ] |
| 36-023 | ToolsDownloadPackageGetV2Api | 下载包查询 | [ ] |
| 36-024 | ToolsDownloadPackageParseV2Api | 解析下载包 | [ ] |
| 36-025 | ToolsMicroAppCreateV30Api | 创建小程序 | [ ] |
| 36-026 | ToolsMicroAppListV30Api | 小程序列表 | [ ] |
| 36-027 | ToolsMicroAppUpdateV30Api | 更新小程序 | [ ] |
| 36-028 | ToolsMicroGameCreateV30Api | 创建小游戏 | [ ] |
| 36-029 | ToolsMicroGameListV30Api | 小游戏列表 | [ ] |
| 36-030 | ToolsMicroGameUpdateV30Api | 更新小游戏 | [ ] |
| 36-031 | ToolsMicroGameConvertWindowGetV30Api | 转化窗口查询 | [ ] |
| 36-032 | ToolsMicroGameConvertWindowUpdateV30Api | 更新转化窗口 | [ ] |
| 36-033 | ToolsWechatAppletCreateV30Api | 创建微信小程序 | [ ] |
| 36-034 | ToolsWechatAppletListV30Api | 微信小程序列表 | [ ] |
| 36-035 | ToolsWechatAppletUpdateV30Api | 更新微信小程序 | [ ] |
| 36-036 | ToolsWechatGameCreateV30Api | 创建微信小游戏 | [ ] |
| 36-037 | ToolsWechatGameListV30Api | 微信小游戏列表 | [ ] |
| 36-038 | ToolQuickAppManagementQuickAppGetV2Api | 快应用查询 | [ ] |

## 37. 其他工具

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 37-001 | ToolsAdPreviewQrcodeGetV30Api | 广告预览二维码 | [ ] |
| 37-002 | ToolsAdminInfoV2Api | 管理员信息 | [ ] |
| 37-003 | ToolsAdvertiserStoreSearchV2Api | 店铺搜索 | [ ] |
| 37-004 | ToolsAssetLinkListV30Api | 资产链接列表 | [ ] |
| 37-005 | ToolsBpAssetManagementShareV30Api | 资产分享 | [ ] |
| 37-006 | ToolsBpAssetManagementShareGetV30Api | 查询分享 | [ ] |
| 37-007 | ToolsBpAssetManagementShareCancelV30Api | 取消分享 | [ ] |
| 37-008 | ToolsEventAllAssetsListV2Api | 所有资产列表 | [ ] |
| 37-009 | ToolsEventAllAssetsDetailV2Api | 资产详情 | [ ] |
| 37-010 | ToolsEventAssetsGetV2Api | 事件资产 | [ ] |
| 37-011 | ToolsEventConvertOptimizedGoalGetV30Api | 转化优化目标 | [ ] |
| 37-012 | ToolsForbiddenLinkGreyGetV30Api | 禁链灰度 | [ ] |
| 37-013 | ToolsGrayGetV30Api | 灰度查询 | [ ] |
| 37-014 | ToolsHotMaterialDeriveListV30Api | 热门素材列表 | [ ] |
| 37-015 | ToolsHotMaterialDeriveGetV30Api | 热门素材查询 | [ ] |
| 37-016 | ToolsHotMaterialDeriveSubmitV30Api | 提交热门素材 | [ ] |
| 37-017 | ToolsHotMaterialDeriveAdoptV30Api | 采纳热门素材 | [ ] |
| 37-018 | ToolsInactiveAdvertiserListV30Api | 不活跃广告主 | [ ] |
| 37-019 | ToolsIsSupportUniversalGetV2Api | 是否支持通用 | [ ] |
| 37-020 | ToolsLogSearchV2Api | 日志搜索 | [ ] |
| 37-021 | ToolsLogSearchDetailGetV2Api | 日志详情 | [ ] |
| 37-022 | ToolsPromotionCardRecommendGetV2Api | 推广卡推荐 | [ ] |
| 37-023 | ToolsPromotionCardRecommendTitleGetV2Api | 推广卡标题 | [ ] |
| 37-024 | ToolsQuotaGetV2Api | 配额查询 | [ ] |
| 37-025 | ToolsVideoCoverSuggestV2Api | 视频封面建议 | [ ] |
| 37-026 | ToolsVideoCheckAvailableAnchorV2Api | 视频可用锚点 | [ ] |
| 37-027 | ToolsPreAuditGetV2Api | 预审查询 | [ ] |
| 37-028 | ToolsPreAuditSendV2Api | 发送预审 | [ ] |
| 37-029 | ToolsPioneerProgramAttachmentUploadV2Api | 先锋计划上传 | [ ] |
| 37-030 | RecommendVideoListV30Api | 推荐视频列表 | [ ] |
| 37-031 | AsyncTaskCreateV2Api | 创建异步任务 | [ ] |
| 37-032 | AsyncTaskGetV2Api | 查询异步任务 | [ ] |
| 37-033 | AsyncTaskDownloadV2Api | 下载异步任务 | [ ] |

## 38. RTA

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 38-001 | ToolsRtaGetV2Api | RTA查询 | [ ] |
| 38-002 | ToolsRtaGetInfoV2Api | RTA信息 | [ ] |
| 38-003 | ToolsRtaGetInfoTmpV2Api | RTA临时信息 | [ ] |
| 38-004 | ToolsRtaStatusUpdateV2Api | RTA状态更新 | [ ] |
| 38-005 | ToolsRtaSetScopeV2Api | RTA范围设置 | [ ] |
| 38-006 | ToolsRtaScopeGetV30Api | RTA范围查询 | [ ] |

## 39. 可玩广告/Rubeex

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 39-001 | ToolsPlayableCreateV2Api | 创建可玩广告 | [ ] |
| 39-002 | ToolsPlayableUploadV2Api | 上传可玩广告 | [ ] |
| 39-003 | ToolsPlayableValidateV2Api | 验证可玩广告 | [ ] |
| 39-004 | ToolsPlayableSaveV2Api | 保存可玩广告 | [ ] |
| 39-005 | ToolsPlayableListGetV2Api | 可玩广告列表 | [ ] |
| 39-006 | ToolsPlayableGrantV2Api | 授权可玩广告 | [ ] |
| 39-007 | ToolsPlayableGrantResultV2Api | 授权结果 | [ ] |
| 39-008 | ToolsPlayableCloudGameListV2Api | 云游戏列表 | [ ] |
| 39-009 | ToolsRubeexVersionGetV2Api | Rubeex版本 | [ ] |
| 39-010 | ToolsRubeexPlayableListV2Api | Rubeex列表 | [ ] |
| 39-011 | ToolsRubeexPlayableAdListV2Api | Rubeex广告列表 | [ ] |
| 39-012 | ToolsRubeexRemarkV2Api | Rubeex备注 | [ ] |

## 40. 原生锚点

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 40-001 | NativeAnchorCreateV30Api | 创建锚点 | [ ] |
| 40-002 | NativeAnchorGetV30Api | 查询锚点 | [ ] |
| 40-003 | NativeAnchorGetDetailV30Api | 锚点详情 | [ ] |
| 40-004 | NativeAnchorUpdateV30Api | 更新锚点 | [ ] |
| 40-005 | NativeAnchorDeleteV30Api | 删除锚点 | [ ] |
| 40-006 | NativeAnchorQrcodePreviewGetV30Api | 锚点二维码预览 | [ ] |

## 41. 穿山甲/流量包

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 41-001 | ToolsUnionFlowPackageCreateV2Api | 创建流量包 | [ ] |
| 41-002 | ToolsUnionFlowPackageGetV2Api | 查询流量包 | [ ] |
| 41-003 | ToolsUnionFlowPackageUpdateV2Api | 更新流量包 | [ ] |
| 41-004 | ToolsUnionFlowPackageDeleteV2Api | 删除流量包 | [ ] |
| 41-005 | ToolsUnionFlowPackageReportV2Api | 流量包报表 | [ ] |
| 41-006 | ToolsUnionFlowPackagePromotionReportV30Api | 推广流量包报表 | [ ] |

## 42. EBP (巨量体验平台)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 42-001 | EbpAdvertiserListV2Api | EBP广告主列表 | [ ] |
| 42-002 | EbpAdvertiserTaskCreateV2Api | 创建任务 | [ ] |
| 42-003 | EbpAdvertiserTaskListV2Api | 任务列表 | [ ] |
| 42-004 | EbpAdvertiserTaskDownloadV2Api | 下载任务 | [ ] |
| 42-005 | EbpLevelGetV2Api | 等级查询 | [ ] |
| 42-006 | ToolsEbpAppListV30Api | EBP应用列表 | [ ] |
| 42-007 | ToolsEbpAppPublishV30Api | 发布应用 | [ ] |
| 42-008 | ToolsEbpAppUpdateV30Api | 更新应用 | [ ] |
| 42-009 | ToolsEbpAppExtendCreateV30Api | 创建扩展 | [ ] |
| 42-010 | ToolsEbpAppExtendListV30Api | 扩展列表 | [ ] |
| 42-011 | ToolsEbpAppExtendUpdateV30Api | 更新扩展 | [ ] |
| 42-012 | ToolsEbpAppGameBookListV30Api | 游戏预约列表 | [ ] |
| 42-013 | ToolsEbpSubjectListV30Api | 主体列表 | [ ] |
| 42-014 | ToolsEbpAssetAuthCancelV30Api | 取消资产授权 | [ ] |
| 42-015 | ToolsEbpAssetAuthListV30Api | 资产授权列表 | [ ] |
| 42-016 | ToolsEbpMaterialAuthCreateV30Api | 创建素材授权 | [ ] |
| 42-017 | ToolsEbpMaterialAuthDeleteV30Api | 删除素材授权 | [ ] |
| 42-018 | ToolsEbpMaterialAuthListV30Api | 素材授权列表 | [ ] |
| 42-019 | ToolsEbpVideoUploadV30Api | 上传视频 | [ ] |
| 42-020 | ToolsEbpVideoDeleteV30Api | 删除视频 | [ ] |
| 42-021 | ToolsEbpVideoUpdateV30Api | 更新视频 | [ ] |
| 42-022 | ToolsEbpVideoAttributesListV30Api | 视频属性列表 | [ ] |
| 42-023 | ToolsEbpMicroAppletCreateV30Api | 创建小程序 | [ ] |
| 42-024 | ToolsEbpMicroAppletListV30Api | 小程序列表 | [ ] |
| 42-025 | ToolsEbpMicroAppletUpdateV30Api | 更新小程序 | [ ] |
| 42-026 | ToolsEbpMicroAppletLinkListV30Api | 小程序链接 | [ ] |
| 42-027 | ToolsEbpMicroGameCreateV30Api | 创建小游戏 | [ ] |
| 42-028 | ToolsEbpMicroGameListV30Api | 小游戏列表 | [ ] |
| 42-029 | ToolsEbpMicroGameUpdateV30Api | 更新小游戏 | [ ] |
| 42-030 | ToolsEbpMicroGameLinkListV30Api | 小游戏链接 | [ ] |
| 42-031 | ToolsEbpWechatAppletCreateV30Api | 微信小程序 | [ ] |
| 42-032 | ToolsEbpWechatAppletListV30Api | 微信小程序列表 | [ ] |
| 42-033 | ToolsEbpWechatAppletUpdateV30Api | 更新微信小程序 | [ ] |
| 42-034 | ToolsEbpWechatGameCreateV30Api | 微信小游戏 | [ ] |
| 42-035 | ToolsEbpWechatGameListV30Api | 微信小游戏列表 | [ ] |
| 42-036 | ToolsEbpWechatGameUpdateV30Api | 更新微信小游戏 | [ ] |

## 43. 品牌广告

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 43-001 | BrandCampaignAddV30Api | 添加品牌Campaign | [ ] |
| 43-002 | BrandCampaignListV30Api | 品牌Campaign列表 | [ ] |
| 43-003 | BrandCampaignGetV30Api | 查询品牌Campaign | [ ] |
| 43-004 | BrandCampaignEditV30Api | 编辑品牌Campaign | [ ] |
| 43-005 | BrandCampaignUpdateV30Api | 更新品牌Campaign | [ ] |
| 43-006 | BrandCampaignSubmitV30Api | 提交品牌Campaign | [ ] |
| 43-007 | BrandCampaignModifyV30Api | 修改品牌Campaign | [ ] |
| 43-008 | BrandCampaignRevokeModifyV30Api | 撤销修改 | [ ] |
| 43-009 | BrandCampaignOperateV30Api | 操作品牌Campaign | [ ] |
| 43-010 | BrandCampaignDeleteV30Api | 删除品牌Campaign | [ ] |
| 43-011 | BrandCampaignRemoveV30Api | 移除品牌Campaign | [ ] |
| 43-012 | BrandAdGetV30Api | 品牌广告查询 | [ ] |
| 43-013 | BrandAdUpdateBaseInfoV30Api | 更新品牌广告基础 | [ ] |
| 43-014 | BrandAdUpdateDeliveryInfoV30Api | 更新投放信息 | [ ] |
| 43-015 | BrandAdDeleteV30Api | 删除品牌广告 | [ ] |
| 43-016 | BrandAdCancelDeleteV30Api | 取消删除 | [ ] |
| 43-017 | BrandCreativeCreateV30Api | 创建品牌创意 | [ ] |
| 43-018 | BrandCreativeGetV30Api | 查询品牌创意 | [ ] |
| 43-019 | BrandCreativeUpdateV30Api | 更新品牌创意 | [ ] |
| 43-020 | BrandCreativeDeleteV30Api | 删除品牌创意 | [ ] |
| 43-021 | BrandMaterialCreateV30Api | 创建品牌物料 | [ ] |
| 43-022 | BrandMaterialListV30Api | 品牌物料列表 | [ ] |
| 43-023 | BrandMaterialUpdateV30Api | 更新品牌物料 | [ ] |
| 43-024 | BrandFileVideoUploadV30Api | 上传品牌视频 | [ ] |
| 43-025 | BrandUploadImageV30Api | 上传品牌图片 | [ ] |
| 43-026 | BrandOrderCreateV30Api | 创建品牌订单 | [ ] |
| 43-027 | BrandOrderListV30Api | 品牌订单列表 | [ ] |
| 43-028 | BrandOrderUpdateV30Api | 更新品牌订单 | [ ] |
| 43-029 | BrandOrderDeleteV30Api | 删除品牌订单 | [ ] |
| 43-030 | BrandOrderCancelDeleteV30Api | 取消删除订单 | [ ] |
| 43-031 | BrandContractGetV30Api | 品牌合同查询 | [ ] |
| 43-032 | BrandPolicyListV30Api | 品牌政策列表 | [ ] |
| 43-033 | BrandQueryStockV30Api | 查询库存 | [ ] |
| 43-034 | BrandQueryStockBalanceV30Api (Tools) | 查询库存余量 | [ ] |
| 43-035 | BrandQueryPublishPriceV30Api (Tools) | 查询刊例价 | [ ] |
| 43-036 | BrandToolCreativePreviewV30Api | 创意预览 | [ ] |
| 43-037 | BrandToolMaterialPreviewV30Api | 物料预览 | [ ] |
| 43-038 | BrandRegionGetV30Api | 品牌地域 | [ ] |
| 43-039 | BrandActionCategoryV30Api | 品牌行为类目 | [ ] |
| 43-040 | BrandAnchorListV30Api | 品牌锚点列表 | [ ] |
| 43-041 | BrandAwemeListV30Api | 品牌抖音号列表 | [ ] |
| 43-042 | BrandCustomAudienceListV30Api | 品牌人群包列表 | [ ] |
| 43-043 | BrandOperationLogQueryV30Api | 品牌操作日志 | [ ] |
| 43-044 | BrandQueryYuntu5aBrandCategoryV30Api | 云图5A品类 | [ ] |
| 43-045 | CdpBrandGetV30Api | CDP品牌查询 | [ ] |

## 44. 商品管理 DPA

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 44-001 | DpaProductCreateV2Api | 创建商品 | [ ] |
| 44-002 | DpaProductDetailGetV2Api | 商品详情 | [ ] |
| 44-003 | DpaProductUpdateV2Api | 更新商品 | [ ] |
| 44-004 | DpaProductDeleteV2Api | 删除商品 | [ ] |
| 44-005 | DpaProductStatusBatchUpdateV2Api | 批量更新状态 | [ ] |
| 44-006 | DpaProductAvailablesV2Api | 可用商品 | [ ] |
| 44-007 | DpaDetailGetV2Api | DPA详情 | [ ] |
| 44-008 | DpaCategoryGetV2Api | DPA类目 | [ ] |
| 44-009 | DpaDictGetV2Api | DPA词典 | [ ] |
| 44-010 | DpaMetaGetV2Api | DPA元信息 | [ ] |
| 44-011 | DpaTemplateGetV2Api | DPA模板 | [ ] |
| 44-012 | DpaVideoGetV2Api | DPA视频 | [ ] |
| 44-013 | DpaAlbumCreateV30Api | 创建专辑 | [ ] |
| 44-014 | DpaAlbumStatusGetV30Api | 专辑状态 | [ ] |
| 44-015 | DpaPlayletAuthGetV2Api | 短剧授权 | [ ] |
| 44-016 | DpaBrandFuzzyGetV30Api | 品牌模糊查询 | [ ] |
| 44-017 | DpaCheckIndexEntryProgressV2Api | 索引进度 | [ ] |
| 44-018 | DpaAssetV2ListV2Api | 资产列表 | [ ] |
| 44-019 | DpaAssetV2DetailReadV2Api | 资产详情 | [ ] |
| 44-020 | DpaAssetsDetailReadV2Api | 资产详情v2 | [ ] |
| 44-021 | DpaClueProductSaveV2Api | 保存线索商品 | [ ] |
| 44-022 | DpaClueProductListV2Api | 线索商品列表 | [ ] |
| 44-023 | DpaClueProductDetailV2Api | 线索商品详情 | [ ] |
| 44-024 | DpaClueProductDeleteV2Api | 删除线索商品 | [ ] |
| 44-025 | DpaEbpProductCreateV30Api | EBP商品创建 | [ ] |
| 44-026 | DpaEbpProductListV30Api | EBP商品列表 | [ ] |
| 44-027 | DpaEbpProductDetailGetV30Api | EBP商品详情 | [ ] |
| 44-028 | DpaEbpProductUpdateV30Api | EBP商品更新 | [ ] |
| 44-029 | DpaEbpProductDeleteV30Api | EBP商品删除 | [ ] |
| 44-030 | DpaEbpProductStatusBatchUpdateV30Api | EBP批量状态 | [ ] |
| 44-031 | DpaEbpCategoryGetV30Api | EBP类目 | [ ] |
| 44-032 | DpaEbpDictGetV30Api | EBP词典 | [ ] |
| 44-033 | DpaEbpMetaGetV30Api | EBP元信息 | [ ] |
| 44-034 | DpaEbpLibraryListV30Api | EBP库列表 | [ ] |
| 44-035 | DpaEbpPlayletAuthGetV30Api | EBP短剧授权 | [ ] |
| 44-036 | DpaEbpClueProductSaveV30Api | EBP线索商品 | [ ] |
| 44-037 | DpaEbpClueProductListV30Api | EBP线索列表 | [ ] |
| 44-038 | DpaEbpClueProductGetV30Api | EBP线索查询 | [ ] |
| 44-039 | DpaEbpClueProductDeleteV30Api | EBP线索删除 | [ ] |

## 45. 安全与审核

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 45-001 | SecurityAuditResultsV30Api | 审核结果 | [ ] |
| 45-002 | SecurityOpenMaterialAuditV30Api | 素材审核 | [ ] |
| 45-003 | SecurityCreateAppealV30Api | 创建申诉 | [ ] |
| 45-004 | SecurityGetConsultResultV30Api | 咨询结果 | [ ] |
| 45-005 | SecurityScoreTotalGetV30Api | 信用分 | [ ] |
| 45-006 | SecurityScoreViolationEventGetV30Api | 违规事件 | [ ] |
| 45-007 | SecurityScoreDisposalInfoGetV30Api | 处置信息 | [ ] |
| 45-008 | OpenMaterialAuditProGetV30Api | 素材专业审核 | [ ] |
| 45-009 | OpenMaterialAuditProSubmitV30Api | 提交专业审核 | [ ] |
| 45-010 | PenaltyTaskGetV30Api | 处罚任务 | [ ] |
| 45-011 | RejectMaterialAiRepairGetV30Api | AI修复查询 | [ ] |
| 45-012 | RejectMaterialAiRepairAcceptTaskCreateV30Api | 创建修复任务 | [ ] |
| 45-013 | RejectMaterialAiRepairAcceptTaskListV30Api | 修复任务列表 | [ ] |
| 45-014 | RejectMaterialAiRepairCrossAccountGetV30Api | 跨账户修复 | [ ] |

## 46. 云图

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 46-001 | YuntuBrandInfoGetV30Api | 品牌信息 | [ ] |
| 46-002 | YuntuAudienceInfoCreateV30Api | 创建受众 | [ ] |
| 46-003 | YuntuAudienceInfoGetV30Api | 查询受众 | [ ] |
| 46-004 | YuntuAudienceInfoDeleteV30Api | 删除受众 | [ ] |
| 46-005 | YuntuAudienceLabelCreateV30Api | 创建标签 | [ ] |
| 46-006 | YuntuAudienceLabelGetV30Api | 查询标签 | [ ] |
| 46-007 | YuntuAudienceLabelDeleteV30Api | 删除标签 | [ ] |

## 47. 店铺/优品/游戏

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 47-001 | ShopBonusCreateV30Api | 创建店铺奖金 | [ ] |
| 47-002 | ShopBonusSuccessGetV30Api | 店铺奖金查询 | [ ] |
| 47-003 | DecorationCouponGetV30Api | 装饰券查询 | [ ] |
| 47-004 | GameAddictionIdGetV30Api | 游戏关键行为ID | [ ] |
| 47-005 | GameplayListV30Api | 游戏玩法列表 | [ ] |
| 47-006 | ServeMarketActiveFuncGetV10Api | 服务市场功能 | [ ] |
| 47-007 | ServeMarketCidVerifyTokenV10Api | CID验证 | [ ] |
| 47-008 | ServeMarketOrderGetV10Api | 服务市场订单 | [ ] |
| 47-009 | SpiTaskGetV2Api | SPI任务查询 | [ ] |

## 48. OC项目 (素材管理)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 48-001 | OcProjectMaterialCreateV30Api | 创建素材 | [ ] |
| 48-002 | OcProjectMaterialGetV30Api | 查询素材 | [ ] |
| 48-003 | OcProjectMaterialDeleteV30Api | 删除素材 | [ ] |
| 48-004 | OcProjectMaterialStatusUpdateV30Api | 更新素材状态 | [ ] |
| 48-005 | OcProjectRejectReasonGetV30Api | 拒绝原因 | [ ] |

## 49. 星推 Stardelivery

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 49-001 | StardeliveryTaskListV30Api | 任务列表 | [ ] |
| 49-002 | StardeliveryTaskDetailV30Api | 任务详情 | [ ] |
| 49-003 | StardeliveryTaskCancelV30Api | 取消任务 | [ ] |
| 49-004 | StardeliveryTaskBudgetUpdateV30Api | 更新预算 | [ ] |
| 49-005 | StardeliveryTaskPostEndTimeUpdateV30Api | 更新结束时间 | [ ] |
| 49-006 | StardeliveryTaskShareV30Api | 分享任务 | [ ] |
| 49-007 | StardeliveryTaskUnshareV30Api | 取消分享 | [ ] |
| 49-008 | StardeliveryTaskShareableListV30Api | 可分享列表 | [ ] |
| 49-009 | StardeliveryTaskSharingListV30Api | 分享中列表 | [ ] |
| 49-010 | StardeliveryTaskAuthorDetailV30Api | 达人详情 | [ ] |
| 49-011 | StardeliveryTaskAuthorVideoDetailV30Api | 达人视频详情 | [ ] |
| 49-012 | StardeliveryTaskAuthorVideoAuditV30Api | 达人视频审核 | [ ] |

## 50. Dou+ (抖+)

| 编号 | SDK方法 | 说明 | 状态 |
|------|---------|------|------|
| 50-001 | DouplusOrderCreateV30Api | 创建Dou+订单 | [ ] |
| 50-002 | DouplusOrderListV30Api | Dou+订单列表 | [ ] |
| 50-003 | DouplusOrderCloseV30Api | 关闭Dou+订单 | [ ] |
| 50-004 | DouplusOrderRenewV30Api | 续费Dou+订单 | [ ] |
| 50-005 | DouplusOrderReportV30Api | Dou+报表 | [ ] |
| 50-006 | DouplusOptionalItemsListV30Api | Dou+可选内容 | [ ] |
| 50-007 | DouplusOptionalTargetsListV30Api | Dou+可选定向 | [ ] |
| 50-008 | DouplusRtaGetInfoV30Api | Dou+ RTA信息 | [ ] |
| 50-009 | DouplusRtaSetScopeV30Api | Dou+ RTA范围 | [ ] |

## 51. 巨量千川 Qianchuan (~160个)

> 千川是独立的电商广告平台，API编号前缀 51-xxx，详见 SDK `Qianchuan*` 系列方法。
> 包含: 广告管理、素材管理、报表、全域推广、直播间、商品、人群包等。
> 因数量较多且为独立平台，此处仅列出子模块概览。

| 子模块 | API数量 | 典型方法 |
|--------|---------|----------|
| 广告管理 | ~25 | QianchuanAdCreateV10Api, QianchuanAdGetV10Api |
| Campaign | ~5 | QianchuanCampaignCreateV10Api |
| 全域推广 | ~25 | QianchuanUniPromotionListV10Api |
| 素材/视频 | ~10 | QianchuanVideoGetV10Api |
| 报表 | ~20 | QianchuanReportAdGetV10Api |
| 直播间 | ~10 | QianchuanTodayLiveRoomGetV10Api |
| 财务 | ~3 | QianchuanFinanceWalletGetV10Api |
| 商品/店铺 | ~8 | QianchuanShopGetV10Api |
| DMP/定向 | ~5 | QianchuanDmpAudiencesGetV10Api |
| 工具 | ~10 | QianchuanToolsEstimateAudienceV10Api |
| 抖音号 | ~15 | QianchuanAwemeOrderCreateV10Api |

## 52. 本地推 Local (~50个)

> 本地生活广告平台，API编号前缀 52-xxx，详见 SDK `Local*` 系列方法。

| 子模块 | API数量 | 典型方法 |
|--------|---------|----------|
| 项目管理 | ~8 | LocalProjectCreateV30Api |
| 单元管理 | ~6 | LocalPromotionCreateV30Api |
| 报表 | ~5 | LocalReportAccountGetV30Api |
| 素材/文件 | ~8 | LocalFileVideoUploadV30Api |
| 门店/商品 | ~5 | LocalPoiGetV30Api |
| IM消息 | ~4 | LocalImSendMsgV30Api |
| 其他 | ~10 | LocalChargeListV30Api |

## 53. 巨量星图 Star (~100个)

> 达人营销平台，API编号前缀 53-xxx，详见 SDK `Star*` 系列方法。

| 子模块 | API数量 | 典型方法 |
|--------|---------|----------|
| 信息查询 | ~5 | StarInfoV2Api |
| 项目管理 | ~5 | StarCreateProjectV2Api |
| 需求管理 | ~15 | StarDemandCreateAssignV2Api |
| 挑战赛 | ~12 | StarChallengeListV2Api |
| 订单管理 | ~12 | StarOrderDetailV2Api |
| MCN管理 | ~15 | StarMcnGetAuthorListV2Api |
| 组件/版权 | ~6 | StarComponentCreateLinkV2Api |
| 报表 | ~8 | StarReportOrderOverviewV2Api |
| 增值服务 | ~8 | StarVasCreateBoostItemGroupV2Api |
| 达人 | ~5 | StarDemanderGetCarBrandListV2Api |

---

统计汇总:
- 巨量引擎核心: 01-50 模块, ~520 个 API
- 巨量千川: 51 模块, ~160 个 API
- 本地推: 52 模块, ~50 个 API
- 巨量星图: 53 模块, ~100 个 API
- 其他/杂项: ~298 个 API
- **总计: 1128 个 API**
