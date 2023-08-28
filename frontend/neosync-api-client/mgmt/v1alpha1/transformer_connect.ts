// @generated by protoc-gen-connect-es v0.13.1 with parameter "target=ts,import_extension=none"
// @generated from file mgmt/v1alpha1/transformer.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetTransformersRequest, GetTransformersResponse } from "./transformer_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service mgmt.v1alpha1.TransformerService
 */
export const TransformerService = {
  typeName: "mgmt.v1alpha1.TransformerService",
  methods: {
    /**
     * @generated from rpc mgmt.v1alpha1.TransformerService.GetTransformers
     */
    getTransformers: {
      name: "GetTransformers",
      I: GetTransformersRequest,
      O: GetTransformersResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
