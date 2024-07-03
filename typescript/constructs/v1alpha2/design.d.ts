/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * Schema for design  in v1Beta1
 */
export interface DesignSchema {
  /**
   * Name of the design
   */
  name: string;
  /**
   * Specifies the version of the schema to which the design  conforms.
   */
  schemaVersion?: string;
  /**
   * Version of the design
   */
  version?: string;
  /**
   * Map of component IDs/names to their corresponding component declarations
   */
  services: {
    [k: string]: {
      id?: string;
      name?: string;
      type?: string;
      apiVersion?: string;
      namespace?: string;
      version?: string;
      model?: string;
      isAnnotation?: boolean;
      labels?: {
        [k: string]: string;
      };
      annotations?: {
        [k: string]: string;
      };
      dependsOn?: string[];
      settings?: {
        [k: string]: unknown;
      };
      traits?: {
        [k: string]: unknown;
      };
      [k: string]: unknown;
    };
  };
  [k: string]: unknown;
}